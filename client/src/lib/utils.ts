export interface getUrlOpts {
  protocol?: string;
  searchParams?: [{ key: string; value: string }];
}

interface getUrlOptsInternal {
  protocol: string;
  searchParams: Map<string, string>;
}

export function getUrl(path: string, opts?: getUrlOpts): string {
  const optsInternal = getUrlDefaultOpts(opts);
  const url = new URL(
    path,
    `${optsInternal.protocol}://${process.env.REACT_APP_OWNGAME_URL}`
  );
  optsInternal.searchParams.forEach(
    (val: string, key: string, _: Map<string, string>) => {
      url.searchParams.append(key, val);
    }
  );
  return url.href;
}

function getUrlDefaultOpts(opts?: getUrlOpts): getUrlOptsInternal {
  if (opts == null) {
    opts = {};
  }
  const searchParams = new Map<string, string>();
  if (opts.searchParams != null) {
    opts.searchParams.forEach(kv => searchParams.set(kv.key, kv.value));
  }
  let optsInternal: getUrlOptsInternal = {
    protocol:
      opts.protocol == null || opts.protocol.length === 0
        ? "https"
        : opts.protocol,
    searchParams: searchParams
  };
  return optsInternal;
}
