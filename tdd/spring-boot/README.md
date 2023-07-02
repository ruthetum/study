
```mermaid
erDiagram
    PRODUCT {
        int id
        string name
        int price
        int discount_policy
    }
    ORDER {
        int id
        int product_id
        int quantity
    }
    PRODUCT ||--o| ORDER : has
```