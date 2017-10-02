---
draft: false
date: "2017-10-02T15:15:50"
tags: ["LoggingDB",
"IdentityDB",
"Crypto",
"MessageBus-A",
"File_Share",
]
title: IdentityProvider
categories: ["everything"]
depmap: [ "graph LR",
"style IdentityProvider fill:#d3460d,stroke:#373b7f,stroke-width:2px",
"IdentityProvider -->|SQL|LoggingDB(\"fa:fa-database LoggingDB\")",
"IdentityProvider -->|SQL|IdentityDB(\"fa:fa-database IdentityDB\")",
"IdentityProvider -->|HTTP|Crypto((\"fa:fa-globe Crypto\"))",
"IdentityProvider -->|WCF|MessageBus-A{\"fa:fa-tasks MessageBus-A\"}",
"IdentityProvider -->|SMB/CIFS/NFS|File_Share[\"fa:fa-files-o File_Share\"]",
]
---
			