docker exec witness-db mysql -uroot -pkdfjjrU64fjK58H -e "ALTER TABLE witness.bsc_transfers ADD payload varchar(24576);"
docker exec witness-db mysql -uroot -pkdfjjrU64fjK58H -e "ALTER TABLE witness.ethereum_transfers ADD payload varchar(24576);"
docker exec witness-db mysql -uroot -pkdfjjrU64fjK58H -e "ALTER TABLE witness.heco_transfers ADD payload varchar(24576);"
docker exec witness-db mysql -uroot -pkdfjjrU64fjK58H -e "ALTER TABLE witness.matic_transfers ADD payload varchar(24576);"
docker exec witness-db mysql -uroot -pkdfjjrU64fjK58H -e "ALTER TABLE witness.polis_transfers ADD payload varchar(24576);"
docker exec witness-db mysql -uroot -pkdfjjrU64fjK58H -e "ALTER TABLE witness.iotex_transfers ADD payload varchar(24576);"
docker exec witness-db mysql -uroot -pkdfjjrU64fjK58H -e "ALTER TABLE witness.iotex_to_heco_transfers ADD payload varchar(24576);"
docker exec witness-db mysql -uroot -pkdfjjrU64fjK58H -e "ALTER TABLE witness.iotex_to_polis_transfers ADD payload varchar(24576);"
docker exec witness-db mysql -uroot -pkdfjjrU64fjK58H -e "ALTER TABLE witness.iotex_to_bsc_transfers ADD payload varchar(24576);"
docker exec witness-db mysql -uroot -pkdfjjrU64fjK58H -e "ALTER TABLE witness.iotex_to_matic_transfers ADD payload varchar(24576);"