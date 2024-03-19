window.BENCHMARK_DATA = {
  "lastUpdate": 1710844291261,
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
          "id": "3bc15cf143c5c8ccfbe50dce3f4ee190c9a6fe02",
          "message": "Reset open files setting",
          "timestamp": "2024-03-19T15:31:59+05:30",
          "tree_id": "bcd3da2f495e1f5e3719d8bae7fb32b00f9ec479",
          "url": "https://github.com/Azure/azure-storage-fuse/commit/3bc15cf143c5c8ccfbe50dce3f4ee190c9a6fe02"
        },
        "date": 1710844290241,
        "tool": "customBiggerIsBetter",
        "benches": [
          {
            "name": "create_1000_files_in_10_threads",
            "value": 120.2353515625,
            "unit": "MiB/s"
          },
          {
            "name": "create_1000_files_in_100_threads",
            "value": 115.5263671875,
            "unit": "MiB/s"
          },
          {
            "name": "create_1l_files_in_20_threads",
            "value": 0.1123046875,
            "unit": "MiB/s"
          }
        ]
      }
    ]
  }
}