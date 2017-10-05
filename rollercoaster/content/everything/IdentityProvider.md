---
draft: false
date: "2017-10-05T14:31:35"
tags: ["IdentityDB",
"Crypto",
"MessageBus-A",
"File_Share",
"LoggingDB",
]
title: IdentityProvider
categories: ["everything"]
depmap: [ "graph LR",
"style IdentityProvider fill:#1cb460,stroke:#602f1f,stroke-width:2px",
"IdentityProvider -->|SQL|IdentityDB(\"fa:fa-database IdentityDB\")",
"IdentityProvider -->|HTTP|Crypto((\"fa:fa-globe Crypto\"))",
"IdentityProvider -->|AMQP|MessageBus-A[\"fa:fa-sitemap MessageBus-A\"]",
"IdentityProvider -->|SMB/CIFS/NFS|File_Share[\"fa:fa-files-o File_Share\"]",
"IdentityProvider -->|SQL|LoggingDB(\"fa:fa-database LoggingDB\")",
]
---
			