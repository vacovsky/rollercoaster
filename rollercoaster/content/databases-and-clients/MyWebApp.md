---
draft: false
date: "2017-10-06T11:03:36"
tags: ["LoggingDB",
"IdentityProvider",
"Service-A",
"ServiceAPI",
]
title: MyWebApp
categories: ["databases-and-clients"]
depmap: [ "graph LR",
"style MyWebApp fill:#76e40a,stroke:#809549,stroke-width:2px",
"MyWebApp>\"fa:fa-arrow-right MyWebApp\"] -->LoggingDB(\"fa:fa-database LoggingDB\")",
"MyWebApp>\"fa:fa-arrow-right MyWebApp\"] -->IdentityProvider((\"fa:fa-globe IdentityProvider\"))",
"MyWebApp>\"fa:fa-arrow-right MyWebApp\"] -->Service-A{\"fa:fa-tasks Service-A\"}",
"MyWebApp>\"fa:fa-arrow-right MyWebApp\"] -->ServiceAPI((\"fa:fa-globe ServiceAPI\"))",
]
---
