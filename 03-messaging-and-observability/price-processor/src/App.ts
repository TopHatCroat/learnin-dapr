import { processPriceEvent } from "./processing/ProcessPriceEvent";
import { savePrice } from "./data/PriceRepository";
import { PriceEvent } from "./data/PriceEvent";
import { server } from "./io/Dapr";

export async function app() {
    const pubSubName = process.env.PUB_SUB_NAME;
    const topicName = process.env.TOPIC_NAME;

    if (!pubSubName) {
        throw Error("PUB_SUB_NAME must be set")
    }

    if (!topicName) {
        throw Error("TOPIC_NAME must be set")
    }

    // Dapr subscription routes orders topic to this route
    server.pubsub.subscribe(pubSubName, topicName, async (data: PriceEvent) => {
        console.log(`Event received ${JSON.stringify(data)}`);

        const result = await processPriceEvent(data);

        await savePrice(result.symbol, result.usd, result.eur)
    });

    console.log(`Starting...`);
    await server.start();
}
