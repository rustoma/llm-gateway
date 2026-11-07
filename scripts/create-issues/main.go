// Command create-issues turns the Markdown tickets in tickets/ into GitHub labels,
// milestones and issues, using the GitHub CLI (gh).
//
// Run it from the repository root:
//
//	go run ./scripts/create-issues/main.go --dry-run   # show what would be created
//	go run ./scripts/create-issues/main.go             # create what is missing
//
// It is safe to run again: labels, milestones and issues that already exist are skipped.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
)

// Ticket is one Markdown file from tickets/.
type Ticket struct {
	Path   string
	Fields map[string]string // header fields: id, title, milestone, labels...
	Body   string            // everything after the header
}

var labelColors = map[string]string{
	"setup": "6B7280", "go": "00ADD8", "ai": "D97706", "learning": "16A34A",
	"api": "0D9488", "streaming": "7C3AED", "testing": "EAB308", "ci": "4F46E5",
	"docker": "2496ED", "config": "64748B", "hands-on": "DC2626",
}

func main() {
	dryRun := flag.Bool("dry-run", false, "only print what would be created")
	flag.Parse()

	tickets, err := loadTickets("tickets")
	if err != nil {
		fail(err)
	}
	fmt.Printf("Found %d tickets\n", len(tickets))

	// 1. Labels. --force updates a label that already exists.
	for _, label := range allLabels(tickets) {
		color, ok := labelColors[label]
		if !ok {
			color = "94A3B8"
		}
		if *dryRun {
			fmt.Println("Would create label:", label)
			continue
		}
		if _, err := gh("label", "create", label, "--color", color, "--force"); err != nil {
			fail(err)
		}
	}

	// 2. Milestones that do not exist yet.
	out, err := gh("api", "repos/{owner}/{repo}/milestones?state=all&per_page=100")
	if err != nil {
		fail(err)
	}
	var milestones []struct{ Title string }
	if err := json.Unmarshal([]byte(out), &milestones); err != nil {
		fail(fmt.Errorf("reading milestones: %w", err))
	}
	have := map[string]bool{}
	for _, m := range milestones {
		have[m.Title] = true
	}
	for _, t := range tickets {
		name := t.Fields["milestone"]
		if have[name] {
			continue
		}
		have[name] = true
		fmt.Println(verb(*dryRun), "milestone:", name)
		if !*dryRun {
			if _, err := gh("api", "repos/{owner}/{repo}/milestones", "-f", "title="+name); err != nil {
				fail(err)
			}
		}
	}

	// 3. Issues, skipping titles that already exist.
	out, err = gh("issue", "list", "--state", "all", "--limit", "1000", "--json", "title")
	if err != nil {
		fail(err)
	}
	var issues []struct{ Title string }
	if err := json.Unmarshal([]byte(out), &issues); err != nil {
		fail(fmt.Errorf("reading issues: %w", err))
	}
	existing := map[string]bool{}
	for _, i := range issues {
		existing[i.Title] = true
	}
	for _, t := range tickets {
		title := fmt.Sprintf("[%s] %s", t.Fields["id"], t.Fields["title"])
		if existing[title] {
			fmt.Println("Skipping (already exists):", title)
			continue
		}
		fmt.Println(verb(*dryRun), "issue:", title)
		if *dryRun {
			continue
		}
		body := fmt.Sprintf("Ticket file: `%s`\n\n%s", filepath.ToSlash(t.Path), t.Body)
		args := []string{"issue", "create", "--title", title, "--body", body, "--milestone", t.Fields["milestone"]}
		for _, label := range splitList(t.Fields["labels"]) {
			args = append(args, "--label", label)
		}
		if _, err := gh(args...); err != nil {
			fail(err)
		}
	}
	fmt.Println("Done.")
}

// loadTickets reads every tickets/*/*.md file, sorted by path (which sorts by ticket number).
func loadTickets(dir string) ([]Ticket, error) {
	paths, err := filepath.Glob(filepath.Join(dir, "*", "*.md"))
	if err != nil {
		return nil, err
	}
	sort.Strings(paths)
	var tickets []Ticket
	for _, path := range paths {
		data, err := os.ReadFile(path)
		if err != nil {
			return nil, err
		}
		parts := strings.SplitN(string(data), "---", 3)
		if len(parts) < 3 {
			return nil, fmt.Errorf("%s: missing --- header", path)
		}
		fields := map[string]string{}
		for _, line := range strings.Split(strings.TrimSpace(parts[1]), "\n") {
			key, value, _ := strings.Cut(line, ":")
			fields[strings.TrimSpace(key)] = strings.TrimSpace(value)
		}
		tickets = append(tickets, Ticket{Path: path, Fields: fields, Body: strings.TrimSpace(parts[2])})
	}
	return tickets, nil
}

// gh runs the GitHub CLI and returns what it printed.
func gh(args ...string) (string, error) {
	cmd := exec.Command("gh", args...)
	var stderr strings.Builder
	cmd.Stderr = &stderr
	out, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("gh %s: %w\n%s", strings.Join(args[:min(3, len(args))], " "), err, stderr.String())
	}
	return string(out), nil
}

func allLabels(tickets []Ticket) []string {
	seen := map[string]bool{}
	for _, t := range tickets {
		for _, label := range splitList(t.Fields["labels"]) {
			seen[label] = true
		}
	}
	labels := make([]string, 0, len(seen))
	for label := range seen {
		labels = append(labels, label)
	}
	sort.Strings(labels)
	return labels
}

func splitList(s string) []string {
	var items []string
	for _, item := range strings.Split(s, ",") {
		if item = strings.TrimSpace(item); item != "" {
			items = append(items, item)
		}
	}
	return items
}

func verb(dryRun bool) string {
	if dryRun {
		return "Would create"
	}
	return "Creating"
}

func fail(err error) {
	fmt.Fprintln(os.Stderr, "error:", err)
	os.Exit(1)
}
