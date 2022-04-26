import { DaprClient, DaprServer } from "dapr-client";

const daprHost = process.env.DAPR_HOST || "http://localhost";
const daprHttpPort = process.env.DAPR_HTTP_PORT || "3501";
const serverHost = process.env.SERVER_HOST || "127.0.0.1";
const serverPort = process.env.SERVER_PORT || "8000";

export const server = new DaprServer(serverHost, serverPort, daprHost, daprHttpPort);
export const client = new DaprClient(daprHost, daprHttpPort);
