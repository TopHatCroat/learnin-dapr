import { DaprServer } from "dapr-client";
import { processPriceEvent } from "./processing/ProcessPriceEvent";
import { savePrice } from "./data/PriceRepository";
import { PriceEvent } from "./data/PriceEvent";

const DAPR_HOST = process.env.DAPR_HOST || "http://localhost";
const DAPR_HTTP_PORT = process.env.DAPR_HTTP_PORT || "3501";
const SERVER_HOST = process.env.SERVER_HOST || "127.0.0.1";
const SERVER_PORT = process.env.SERVER_PORT || "5001";

const server = new DaprServer(SERVER_HOST, SERVER_PORT, DAPR_HOST, DAPR_HTTP_PORT);

export async function app() {
    // Dapr subscription routes orders topic to this route
    server.pubsub.subscribe("price_pub_sub", "orders", async (data: PriceEvent) => {
        const result = await processPriceEvent(data);

        await savePrice(result.symbol, result.usd, result.eur)
    });

    await server.start();
}
