const prefix: string = "owngame-";

export const LOCAL_STORAGE_KEY_WS_TOKEN = "ws-token";

export class localStorage {
  static get(key: string): string | null {
    const pKey = this.prefixKey(key, prefix);
    const value = window.localStorage.getItem(pKey);
    return value;
  }

  static set(key: string, value: string): void {
    const pKey = localStorage.prefixKey(key, prefix);
    window.localStorage.setItem(pKey, value);
  }

  static remove(key: string): void {
    const pKey = localStorage.prefixKey(key, prefix);
    window.localStorage.removeItem(pKey);
  }

  private static prefixKey(key: string, prefix: string): string {
    return prefix + key;
  }
}
