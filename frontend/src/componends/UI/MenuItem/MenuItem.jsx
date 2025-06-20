import React from 'react';
import ProductCard from "../ProductCard/ProductCard";
import cl from "./MenuItem.module.css"
import ProductService from "../../../api/ProductService";

const MenuItem = ({products, setProducts}) => {
    const handleDelete = async (id) => {
        console.log(id);
        try {
            const res = await ProductService.deleteProduct(id, products);
            console.log(res.data);
            setProducts(products.filter(p => p.id !== id));

        } catch (e) {
            if (e.status === 401) {
                alert("only admin 🙉");
            }
            if (e.status === 500) {
                alert("something went wrong");
            }
            console.log(e.status);
        }
    };
    return (
        <>
            {(products === null || products.length === 0) ? (
                <div className={cl.noProducts}>
                    <h1>No products found.</h1>
                </div>
            ) : (
                <div className={cl.grid}>
                    {products.map(product => (
                        <div key={product.id} className={cl.cardWrapper}>
                            <ProductCard
                                title={product.name}
                                price={product.price}
                                description={product.description}
                                image={product.image}
                                onDelete={() => handleDelete(product.id)}
                            />
                        </div>
                    ))}
                </div>
            )}
        </>
    );
};

export default MenuItem;