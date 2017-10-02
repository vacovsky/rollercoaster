---
title: "{{ replace .TranslationBaseName "-" " " | title }}"
date: {{ .Date }}
draft: false
categories: ["Dependency Map"]
depmap: ["graph TD",
    "A[Christmas] -->|Get money| B(Go shopping)",
    "B --> C{Let me think}",
    "C -->|One| D[Laptop]",
    "C -->|Two| E[iPhone]",
    "C -->|Three| F[Car]",
    "D --> A{who knows}",
    ]
---
