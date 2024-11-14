import { useLocalStorage } from "@vueuse/core";

export const cartToKlarnaPayload = function () {

    const cart = useLocalStorage('cart', [], {});
    const items = [];

    let order_tax_amount = 0;
    let total_tax_amount = 0;

    cart.value.map((item)=>{
        const price = item.product.price;
        const taxCents = price * 0.19;
        const totalCents = price + taxCents;
        const qty = item.qty;

        items.push({
            "type": "physical",
            "reference": "19-402",
            "name": "T-shirt",
            "quantity": qty,
            "unit_price": totalCents,         // 100.00 EUR
            "tax_rate": taxCents,            // 20% tax rate
            "total_amount": qty * totalCents,       // 100.00 EUR total for the item
            "total_tax_amount": qty * taxCents     // 19.00 EUR total tax
        });
        order_tax_amount+= qty * totalCents;
        total_tax_amount += qty * taxCents;
    });

    const tmp = {
        "purchase_country": "DE",
        "purchase_currency": "EUR",
        "locale": "de-DE",
        "order_amount": order_tax_amount,          // 100.00 EUR in minor units
        "order_tax_amount": total_tax_amount,       // 19.00 EUR tax in minor units
        "order_lines": items
    };

    console.log(tmp);

    return tmp;
}