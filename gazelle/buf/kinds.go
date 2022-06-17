package buf

import "github.com/bazelbuild/bazel-gazelle/rule"

const (
	lintRuleKind             = "buf_lint_test"
	breakingRuleKind         = "buf_breaking_test"
	dependenciesRepoRuleKind = "buf_dependencies"
)

var bufKinds = map[string]rule.KindInfo{
	lintRuleKind: {
		MatchAttrs: []string{"targets"},
		NonEmptyAttrs: map[string]bool{
			"targets": true,
		},
		MergeableAttrs: map[string]bool{
			"config": true,
		},
	},
	breakingRuleKind: {
		MatchAttrs: []string{"targets"},
		NonEmptyAttrs: map[string]bool{
			"targets": true,
		},
		MergeableAttrs: map[string]bool{
			"against":              true,
			"exclude_imports":      true,
			"limit_to_input_files": true,
		},
		ResolveAttrs: map[string]bool{
			"targets": true,
		},
	},
	dependenciesRepoRuleKind: {
		MatchAttrs: []string{"name"},
		NonEmptyAttrs: map[string]bool{
			"modules": true,
		},
		MergeableAttrs: map[string]bool{
			"modules": true,
		},
	},
}

var bufLoads = []rule.LoadInfo{
	{
		Name: "@rules_buf//buf:defs.bzl",
		Symbols: []string{
			lintRuleKind,
			breakingRuleKind,
			dependenciesRepoRuleKind,
		},
	},
}

func (*bufLang) Kinds() map[string]rule.KindInfo { return bufKinds }
func (*bufLang) Loads() []rule.LoadInfo          { return bufLoads }
