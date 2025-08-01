.PHONY: apply-patches apply-sed-changes update-pkg-cache

update-pkg-cache:
	@echo "Updating new package version..."
	GOPROXY=https://proxy.golang.org GO111MODULE=on \
	go get github.com/talon-one/talon_go/v10

apply-patches:
	@echo "Applying gopatch patches..."
	gopatch -p patches/condition.patch model_rule.go
	gopatch -p patches/condition.patch model_generate_rule_title_rule.go
	gopatch -p patches/effects.patch model_event.go
	gopatch -p patches/effects.patch model_rule.go
	gopatch -p patches/effects.patch model_generate_rule_title_rule.go
	gopatch -p patches/expr.patch model_template_def.go
	gopatch -p patches/expr.patch model_new_template_def.go
	gopatch -p patches/expression.patch model_binding.go
	gopatch -p patches/expression.patch model_application_cif_expression.go
	gopatch -p patches/expression.patch model_new_application_cif_expression.go

apply-sed-changes:
	@echo "Executing sed command..."
	@sed -i -f patches/effects.sed docs/Event.md
	@sed -i -f patches/effects.sed docs/GenerateRuleTitleRule.md
	@sed -i -f patches/effects.sed docs/Rule.md
	@sed -i -f patches/expr.sed docs/NewTemplateDef.md
	@sed -i -f patches/expr.sed docs/TemplateDef.md
	@sed -i -f patches/expression.sed docs/ApplicationCifExpression.md
	@sed -i -f patches/expression.sed docs/Binding.md
	@sed -i -f patches/expression.sed docs/NewApplicationCifExpression.md


apply-git-patches:
	git apply patches/git/Binding.patch
	git apply patches/git/Event.patch
	git apply patches/git/GenerateItemFilterDescription.patch
	git apply patches/git/GenerateRuleTitleRule.patch
	git apply patches/git/NewApplicationCIFExpression.patch
	git apply patches/git/NewTemplateDef.patch
	git apply patches/git/Rule.patch
	git apply patches/git/TemplateDef.patch
