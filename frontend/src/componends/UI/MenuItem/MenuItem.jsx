import React from 'react';
import ProductCard from "../ProductCard/ProductCard";
import cl from "./MenuItem.module.css"

const MenuItem = ({products, setProducts}) => {
    const handleDelete = (id) => {
        setProducts(products.filter(p => p.id !== id));
    };
    return (
        <div className={cl.grid}>
            {products.map(product => (
                <div key={product.id} className={cl.cardWrapper}>
                    <ProductCard
                        title={product.title}
                        price={product.price}
                        description={product.description}
                        image={product.image}
                        onDelete={() => handleDelete(product.id)}
                    />
                </div>
            ))}
        </div>
    );
};

export default MenuItem;