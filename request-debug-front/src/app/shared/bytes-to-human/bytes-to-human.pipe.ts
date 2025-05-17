import {Pipe, PipeTransform} from '@angular/core';

@Pipe({
  name: 'bytesToHuman',
  pure: false,
  standalone: true,
})
export class BytesToHumanPipe implements PipeTransform {
  transform(bytes: number, precision = 2): string {
    if (isNaN(parseFloat(String(bytes))) || !isFinite(bytes)) {
      return '-';
    }

    const units = ['b', 'kB', 'MB', 'GB', 'TB', 'PB'];
    let unit = 0;

    while (bytes >= 1024) {
      bytes /= 1024;
      unit++;
    }

    return bytes.toFixed(+precision) + ' ' + units[unit];
  }
}
