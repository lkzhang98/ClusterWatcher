# 默认执行 all 目标
.DEFAULT_GOAL := all

# =================================================================
# 定义 Makefile all 伪目标，执行 `make` 时， 会默认执行 all 伪目标

.PHONY: all
all: gen.add-copyright go.format go.build

# =================================================================
# I