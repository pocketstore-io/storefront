import { useLocalStorage } from "@vueuse/core";

const cart = useLocalStorage('cart', [], {})

export const cartToPurchaseUnits = function () {
    let total = 0;
    cart.value.map((item)=>{
        total += item.qty * item.product.price;
    });

    return [{
        amount: {
            currency_code: 'EUR',
            value: total,
        }
    }];
}