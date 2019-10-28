// TODO: DRY!
declare const process: {
  env: {
    REACT_APP_OWNGAME_URL: string;
  };
};

export class WS {
  private url: string = "";
  private conn: WebSocket | null = null;

  async open(path: string): Promise<void> {
    return new Promise((resolve, reject) => {
      this.url = WS.composeURL(path);
      this.conn = new WebSocket(this.url);
      this.conn.onerror = e => {
        reject("failed to create a websocket connection with " + this.url);
      };
      this.conn.onopen = e => {
        resolve();
      };
    });
  }

  private static composeURL(path: string): string {
    return `ws://${process.env.REACT_APP_OWNGAME_URL}/${path}`;
  }
}
