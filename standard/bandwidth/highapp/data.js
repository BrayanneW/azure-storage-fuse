window.BENCHMARK_DATA = {
  "lastUpdate": 1719989712117,
  "repoUrl": "https://github.com/Azure/azure-storage-fuse",
  "entries": {
    "Benchmark": [
      {
        "commit": {
          "author": {
            "email": "vibhansa@microsoft.com",
            "name": "vibhansa",
            "username": "vibhansa-msft"
          },
          "committer": {
            "email": "vibhansa@microsoft.com",
            "name": "vibhansa",
            "username": "vibhansa-msft"
          },
          "distinct": true,
          "id": "65a677429517c28a75a8ab6a3051311f44ba96aa",
          "message": "Revert back to parallel runs",
          "timestamp": "2024-07-02T20:58:56-07:00",
          "tree_id": "6f63f586f70486b7a8ef5ccd4a64d7862cb7715c",
          "url": "https://github.com/Azure/azure-storage-fuse/commit/65a677429517c28a75a8ab6a3051311f44ba96aa"
        },
        "date": 1719989711895,
        "tool": "customBiggerIsBetter",
        "benches": [
          {
            "name": "read_10_20GB_file",
            "value": 24.022892653456438,
            "unit": "GiB/s"
          },
          {
            "name": "create_10_20GB_file",
            "value": 30.589700134591034,
            "unit": "GiB/s"
          }
        ]
      }
    ]
  }
}