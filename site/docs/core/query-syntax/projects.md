---
title: "Runs Query Syntax"
sub_link: "query-syntax/runs"
meta_title: "Polyaxon Experiments Query Syntax Specification - Polyaxon References"
meta_description: "Polyaxon Experiments Query Syntax Specification."
visibility: public
status: published
tags:
    - api
    - reference
    - polyaxon
    - query
    - syntax
    - search
    - experiment
    - tracking
    - insights
sidebar: "core"
---

# Searching Projects

## Query

field                           | condition
--------------------------------|------------------
`id`                            | [value condition](/docs/core/query-syntax/#query-with-value-condition)
`name`                          | [string condition](/docs/core/query-syntax/#query-with-string-condition)
`description`                   | [string condition](/docs/core/query-syntax/#query-with-string-condition)
`created_at`                    | [datetime condition](/docs/core/query-syntax/#query-with-datetime-condition)
`updated_at`                    | [datetime condition](/docs/core/query-syntax/#query-with-datetime-condition)
`user`                          | [value condition](/docs/core/query-syntax/#query-with-value-condition)
`tags`                          | [value condition](/docs/core/query-syntax/#query-with-value-condition)


## Sort

field:

 * `created_at`
 * `updated_at`
 * `inputs.*`
 * `outputs.*`
 * `metrics.*`
