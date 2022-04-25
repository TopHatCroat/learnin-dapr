import { PriceEvent } from "../data/PriceEvent";

export function processPriceEvent(data: PriceEvent): Promise<PriceEvent> {
    console.log("Subscriber received: " + JSON.stringify(data))

    return Promise.resolve(data)
}
