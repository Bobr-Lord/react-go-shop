import React, {useContext, useEffect} from 'react';
import Loader from "../UI/Loader/Loader";
import MenuItem from "../UI/MenuItem/MenuItem";
import {useFetching} from "../../hooks/useFetching";
import ProductService from "../../api/ProductService";
import {AuthContext} from "../../context";

const MenuItemLoader = ({products, setProducts}) => {
    const {isLoggedIn} = useContext(AuthContext);
    const [fetchProduct, isLoading, ProductError] = useFetching(async () => {
        let FetchProducts;
        if (isLoggedIn) {
            FetchProducts = await ProductService.getProductsPrivate();
            console.log(FetchProducts);
        } else{
            FetchProducts = await ProductService.getProducts();
        }
        setProducts(FetchProducts.data.products);
        console.log(FetchProducts);
    })
    useEffect(() => {
        fetchProduct();
    }, [])


    return (
        <div>
            {ProductError && <h1>Произошла ошибка</h1>}
            {!ProductError && (
                (isLoading)
                ? <div style={{display: "flex", justifyContent: "center", marginTop: 50}}> <Loader/> </div>
                : <MenuItem products={products} setProducts={setProducts} />
            )}
        </div>
    );
};

export default MenuItemLoader;