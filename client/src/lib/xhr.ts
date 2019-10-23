declare const process: {
  env: {
    OWNGAME_URL: string;
  };
};

export interface XHRResponse {
  status: number;
  responseText?: string;
  responseType?: XMLHttpRequestResponseType;
}

export enum XHRRequestContentType {
  JSON,
  Unknown
}

export interface XHRRequestOptions {
  method: string;
  path: string;
  contentType?: XHRRequestContentType;
  body?: any;
}

export class XHRRequest {
  static async send(opts: XHRRequestOptions): Promise<XHRResponse> {
    return new Promise((resolve, reject) => {
      if (opts == null) {
        throw new Error("no XHRRequestOptions passed");
      }
      opts.contentType =
        opts.contentType == null
          ? XHRRequestContentType.JSON
          : opts.contentType;
      let xhr = new XMLHttpRequest();
      const rejectFn = (xhr: XMLHttpRequest, reject: (reason: any) => void) => {
        reject({
          status: xhr.status,
          statusText: xhr.statusText
        });
      };
      // onload
      xhr.onload = function() {
        if (this.status >= 200 && this.status < 300) {
          let resp: XHRResponse = {
            status: this.status,
            responseText: this.responseText,
            responseType: this.responseType
          };
          resolve(resp);
          return;
        }
        rejectFn(this, reject);
      };
      //onerror
      xhr.onerror = function() {
        rejectFn(this, reject);
      };
      // compose and send
      const url = XHRRequest.getUrl(opts.path);
      xhr.open(opts.method, url);
      const headers = new XHRRequestHeaders().addContentType(opts.contentType);
      headers.forEach((key, string) => {
        console.log(`${key}: ${string}`);
        xhr.setRequestHeader(key, string);
      });
      let bodyEnc = XHRRequest.encodeBody(opts.body, opts.contentType);
      xhr.send(bodyEnc);
    });
  }

  private static getUrl(path: string): string {
    console.log(process.env.OWNGAME_URL);
    return `http://${process.env.OWNGAME_URL}/${path}`;
  }

  private static encodeBody(
    body: any,
    contentType?: XHRRequestContentType
  ): string | null {
    if (body == null) {
      return null;
    }
    switch (contentType) {
      case XHRRequestContentType.JSON:
        return JSON.stringify(body);
      default:
        return null;
    }
  }
}

class XHRRequestHeaders {
  private headers: Headers;

  constructor() {
    this.headers = new Headers();
  }

  addContentType(contentType: XHRRequestContentType): XHRRequestHeaders {
    switch (contentType) {
      case XHRRequestContentType.JSON:
        this.headers.append("Content-Type", "application/json");
        break;
      default:
        break;
    }
    return this;
  }

  forEach(callback: (key: string, value: string) => void) {
    for (const keyVal of this.headers) {
      callback(keyVal[0], keyVal[1]);
    }
  }
}
