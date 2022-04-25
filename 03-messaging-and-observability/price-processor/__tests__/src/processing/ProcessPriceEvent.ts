import { PriceEvent } from "../../../src/data/PriceEvent";
import { processPriceEvent } from "../../../src/processing/ProcessPriceEvent";

it("should process event", async () => {
  const data: PriceEvent = {
    symbol: "BTC",
    eur: 38034.537,
    usd: 41467.0555
  }

  const result =  await processPriceEvent(data);

  expect(result).toEqual(data);
})
