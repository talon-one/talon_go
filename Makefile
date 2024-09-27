update-pkg-cache:
	GOPROXY=https://proxy.golang.org GO111MODULE=on \
	go get github.com/talon-one/talon_go/v8

apply-patches:
	gopatch -p patches/effects.patch model_event.go
	gopatch -p patches/effects.patch model_rule.go
	gopatch -p patches/effects.patch model_generate_rule_title_rule.go
	gopatch -p patches/expression.patch model_binding.go
	gopatch -p patches/expression.patch model_application_cif_expression.go
	gopatch -p patches/expression.patch model_new_application_cif_expression.go
	gopatch -p patches/expr.patch model_template_def.go
	gopatch -p patches/expr.patch model_new_template_def.go
	gopatch -p patches/condition.patch model_rule.go
	gopatch -p patches/condition.patch model_generate_rule_title_rule.go
