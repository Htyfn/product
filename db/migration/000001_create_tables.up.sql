CREATE TABLE IF NOT EXISTS product (
    Id          INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    seller      INT NOT NULL,
    price       MONEY NOT NULL,
    curr        INT NOT NULL,
    createDate  DATE,
    updateDate  DATE
);
CREATE INDEX IF NOT EXISTS product_idx_sellerId ON product(seller);

CREATE TABLE IF NOT EXISTS productAttrStr (
    ProductId   INT references product(Id),
    Key			VARCHAR(500),
    Value	    VARCHAR(4000)
);
CREATE INDEX IF NOT EXISTS ProductAttrStr_idx_productId ON product(Id);
