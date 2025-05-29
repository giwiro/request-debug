const localStoragePolyfill = {
  _data: {} as Record<string, string>,
  setItem: function (id: string, val: string) {
    this._data[id] = val;
  },
  getItem: function (id: string) {
    return Object.prototype.hasOwnProperty.call(this._data, id)
      ? this._data[id]
      : null;
  },
  removeItem: function (id: string) {
    delete this._data[id];
  },
  clear: function () {
    this._data = {};
  },
  length: function () {
    return Object.keys(this._data).length;
  },
};

export const localStorage = window.localStorage
  ? window.localStorage
  : localStoragePolyfill;
