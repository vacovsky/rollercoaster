---
draft: false
date: "2017-10-02T11:28:22"
tags: ["IdentityProvider",
"Service-A",
"ServiceAPI",
]
title: MyWebApp
categories: ["services-and-apis"]
depmap: [ "graph LR",
"style MyWebApp fill:#0118d9,stroke:#2e962d,stroke-width:2px",
"MyWebApp[\"fa:fa-arrow-right MyWebApp\"] -->|HTTP|IdentityProvider((\"fa:fa-globe IdentityProvider\"))",
"MyWebApp[\"fa:fa-arrow-right MyWebApp\"] -->|WCF|Service-A{\"fa:fa-tasks Service-A\"}",
"MyWebApp[\"fa:fa-arrow-right MyWebApp\"] -->|HTTP|ServiceAPI((\"fa:fa-globe ServiceAPI\"))",
]
---
			