import { getUrl } from "./utils";

// TODO: DRY!
declare const process: {
  env: {
    REACT_APP_OWNGAME_URL: string;
  };
};

export class WS {
  private url: string = "";
  private conn: WebSocket | null = null;

  async open(
    path: string,
    onmessage: (data: MessageEvent) => void,
    params?: [{ key: string; value: string }],
    onerror?: (e: ErrorEvent) => void
  ): Promise<void> {
    return new Promise((resolve, reject) => {
      this.url = getUrl(path, {
        protocol: "ws",
        searchParams: params
      });
      this.conn = new WebSocket(this.url);
      this.conn.onerror = e => {
        reject("failed to create a websocket connection with " + this.url);
      };
      this.conn.onmessage = e => onmessage(e);
      this.conn.onopen = e => {
        resolve();
      };
    });
  }

  private static composeURL(path: string): string {
    return `https://${process.env.REACT_APP_OWNGAME_URL}/${path}`;
  }
}
