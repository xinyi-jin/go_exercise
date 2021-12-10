package alg748

import "testing"

func Test_shortestCompletingWord(t *testing.T) {
	type args struct {
		licensePlate string
		words        []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Test_shortestCompletingWord_01", args{"1s3 PSt", []string{"step", "steps", "stripe", "stepple"}}, "steps"},
		{"Test_shortestCompletingWord_02", args{"1s3 456", []string{"looks", "pest", "stew", "show"}}, "pest"},
		{"Test_shortestCompletingWord_03", args{"Ah71752", []string{"suggest", "letter", "of", "husband", "easy", "education", "drug", "prevent", "writer", "old"}}, "husband"},
		{"Test_shortestCompletingWord_04", args{"OgEu755", []string{"enough", "these", "play", "wide", "wonder", "box", "arrive", "money", "tax", "thus"}}, "enough"},
		{"Test_shortestCompletingWord_05", args{"iMSlpe4", []string{"claim", "consumer", "student", "camera", "public", "never", "wonder", "simple", "thought", "use"}}, "simple"},
		{"Test_shortestCompletingWord_05", args{"AN87005", []string{"participant", "individual", "start", "exist", "above", "already", "easy", "attack", "player", "important"}}, "important"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shortestCompletingWord(tt.args.licensePlate, tt.args.words); got != tt.want {
				t.Errorf("%q. shortestCompletingWord() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
