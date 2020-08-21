// This file can be replaced during build by using the `fileReplacements` array.
// `ng build --prod` replaces `environment.ts` with `environment.prod.ts`.
// The list of file replacements can be found in `angular.json`.

export const environment = {
  production: false,
  apiRequestsUrl: 'http://localhost:7005/goscope/api/requests',
  apiSearchRequestsUrl: 'http://localhost:7005/goscope/api/search/requests',
  apiLogsUrl: 'http://localhost:7005/goscope/api/logs',
  apiSearchLogsUrl: 'http://localhost:7005/goscope/api/search/logs',
  apiSysInfoUrl: 'http://localhost:7005/goscope/api/info',
  apiApplicationNameUrl: 'http://localhost:7005/goscope/api/application-name'
};

export const environment = {
  production: true,
  apiRequestsUrl: '/goscope/api/requests',
  apiSearchRequestsUrl: '/goscope/api/search/requests',
  apiLogsUrl: '/goscope/api/logs',
  apiSearchLogsUrl: '/goscope/api/search/logs',
  apiSysInfoUrl: '/goscope/api/info',
  apiApplicationNameUrl: '/goscope/api/application-name'
};
/*
 * For easier debugging in development mode, you can import the following file
 * to ignore zone related error stack frames such as `zone.run`, `zoneDelegate.invokeTask`.
 *
 * This import should be commented out in production mode because it will have a negative impact
 * on performance if an error is thrown.
 */
// import 'zone.js/dist/zone-error';  // Included with Angular CLI.
