# pgxmock-examples
Examples to mock [pgx](https://github.com/jackc/pgx)

## Overview
All examples are based on [pgxmock](https://github.com/pashagolub/pgxmock) and [pgxpoolmock](https://github.com/driftprogramming/pgxpoolmock), thanks for the great works from the authors of these repositories!  

This repo add some mock examples based on different structure level.

|                   |       pgxmock      |     pgxmock/v2     |     pgxpoolmock    |
| :----:            |       :----:       |       :----:       |       :----:       |
| pgx version       |         v4         |         v5         |         v4         |
| pgx.Conn mock     | :heavy_check_mark: | :heavy_check_mark: |         :x:        |
| pgxpool.Pool mock | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: |
| pgxpool.Conn mock | :heavy_check_mark: | :heavy_check_mark: |         :x:        |

## Usage
- Mock level `pgx.Conn`, see [pgxconn.go](./pgxconn.go) and [pgxconn_test.go](./pgxconn_test.go)
- Mock level `pgxpool.Pool`, see [pgxpool.go](./pgxpool.go) and [pgxpool_test.go](./pgxpool_test.go)
- Mock level `pgxpool.Conn`, see [pgxpoolconn.go](./pgxpoolconn.go) and [pgxpoolconn_test.go](./pgxpoolconn_test.go)