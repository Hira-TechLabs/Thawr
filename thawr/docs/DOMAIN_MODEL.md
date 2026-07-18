# Thawr Domain Model

Organization
│
├── Organization Services
│   ├── Identity
│   ├── AI
│   ├── Networking
│   ├── Monitoring
│   └── Backup
│
└── Applications
    │
    ├── Application A
    │   ├── Compute
    │   ├── Database
    │   ├── Cache
    │   ├── Storage
    │   ├── Secrets
    │   └── Routes
    │
    ├── Application B
    └── Application N




## Rules

- Every Application belongs to exactly one Organization.
- Organization services may be shared across applications.
- Application resources are isolated by default.
- No Application can exist without an Organization.
