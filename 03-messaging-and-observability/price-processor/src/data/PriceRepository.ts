import { DaprClient, DaprServer } from "dapr-client";
import { v4 as uuid } from 'uuid';
import { client } from "../io/Dapr";

const STORE_NAME = "price"

export async function savePrice(currency: string, usdPrice: number, eurPrice: number) {
    const price = {
        currency,
        usdPrice,
        eurPrice
    };

    await client.state.save(STORE_NAME, [{ key: uuid(), value: price }]);
}
