package crontab

import "testing"

// Tests the parsing of valid crontab entries.
func TestParseValid(t *testing.T) {
	tests := []struct{ input, command, dump string }{
		// Extraneous whitespace.
		{"  1-3      6,18  1 */6 0-2     /bin/true", "/bin/true", "1,2,3 6,18 1 1,7 0,1,2"},

		// Duplicate values.
		{"9,8,7,1,2,3,2,1,7,8,9 0 1 1 0 /bin/true", "/bin/true", "1,2,3,7,8,9 0 1 1 0"},

		// Overlapping ranges.
		{"3-7,5-9 0 1 1 0 /bin/true", "/bin/true", "3,4,5,6,7,8,9 0 1 1 0"},

		// One range contains another.
		{"0 5-12,6-11 1 1 0 /bin/true", "/bin/true", "0 5,6,7,8,9,10,11,12 1 1 0"},

		// Steps.
		{"*/15 */7 5-20/3 1 0 /bin/true", "/bin/true", "0,15,30,45 0,7,14,21 5,8,11,14,17,20 1 0"},

		// Step with implied range.
		{"30/9 0 1 1 0 /bin/true", "/bin/true", "30,39,48,57 0 1 1 0"},

		// Number of days in the month is independent of the month.
		{"*/15 0 * 2 0 /bin/true", "/bin/true", "0,15,30,45 0 1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31 2 0"},

		// A variety of constructs.
		{"1-13/2 * */7 1,7 * /x/y/z -t foo --x=bar", "/x/y/z -t foo --x=bar", "1,3,5,7,9,11,13 0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23 1,8,15,22,29 1,7 0,1,2,3,4,5,6"},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			job, err := ParseJob(tt.input)
			switch {
			case err != nil:
				t.Fatalf("parse failed: %s", err)
			case job.Command != tt.command:
				t.Fatalf("job.Command is %q, want %q", job.Command, tt.command)
			case job.dump() != tt.dump:
				t.Fatalf("job.dump() returned %q, want %q", job.dump(), tt.dump)
			}
		})
	}
}

// Tests the parsing of invalid crontab entries.
func TestParseInvalid(t *testing.T) {
	tests := []struct{ input string }{
		// Missing time field.
		{"* 0 * 1 /bin/true foo bar baz"},

		// Empty command.
		{"0 0 1 1 0         "},

		// Extra step.
		{"*/1/2 0 1 1 0 /bin/true"},

		// Zero step size.
		{"*/0 0 1 1 0 /bin/true"},

		// Negative step size.
		{"*/-2 0 1 1 0 /bin/true"},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if _, err := ParseJob(tt.input); err == nil {
				t.Fatal("expected a parse error")
			}
		})
	}
}
