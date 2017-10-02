---
draft: false
date: "2017-10-01T22:40:26"
tags: ["MessageBus-B",
"MessageBus-A",
"Identity",
]
title: ServiceA
categories: ["services-and-apis"]
depmap: [ "graph LR",
"style ServiceA fill:#3cde8c,stroke:#552523,stroke-width:2px",
"ServiceA -->|WCF|MessageBus-B{\"fa:fa-tasks MessageBus-B\"}",
"ServiceA -->|WCF|MessageBus-A{\"fa:fa-tasks MessageBus-A\"}",
"ServiceA -->|HTTP|Identity((\"fa:fa-globe Identity\"))",
]
---
			