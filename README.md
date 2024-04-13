
## Acey Ducey

A simple port of the Creative Computing game.
https://www.atariarchives.org/basicgames/showpage.php?page=2



```mermaid
flowchart TD
    
    A[Introduction] --> B[Your money remaining]
    B --> C[Draw card]
    C --> D[Display cards]
    D --> E[Get bet]
    E --> F[Draw card - final]
    F --> Y[Try Again?]
    Y --> |Yes| B
    Y --> |No| Z
    Z[END]


```
