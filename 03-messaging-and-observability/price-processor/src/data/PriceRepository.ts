import { DaprClient, DaprServer } from "dapr-client";
import { v4 as uuid } from 'uuid';

const DAPR_HOST = process.env.DAPR_HOST || "http://localhost";
const DAPR_HTTP_PORT = process.env.DAPR_HTTP_PORT || "3501";

const STORE_NAME = "price"

const client = new DaprClient(DAPR_HOST, DAPR_HTTP_PORT);

export async function savePrice(currency: string, usdPrice: number, eurPrice: number) {
    const price = {
        currency,
        usdPrice,
        eurPrice
    };

    await client.state.save(STORE_NAME, [{ key: uuid(), value: price }]);
}
