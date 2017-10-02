---
draft: false
date: "2017-10-01T22:40:26"
tags: ["Identity",
"MessageBus-B",
"MessageBus-C",
]
title: ServiceB
categories: ["services-and-apis"]
depmap: [ "graph LR",
"style ServiceB fill:#76e40a,stroke:#809549,stroke-width:2px",
"ServiceB -->|HTTP|Identity((\"fa:fa-globe Identity\"))",
"ServiceB -->|WCF|MessageBus-B{\"fa:fa-tasks MessageBus-B\"}",
"ServiceB -->|WCF|MessageBus-C{\"fa:fa-tasks MessageBus-C\"}",
]
---
			