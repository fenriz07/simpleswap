# simpleswap
**simpleswap** is a blockchain built using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli).

## Get started

### 1 Creating the image

```
docker build . -t simpleswap

```
### 2 Running the container

```
docker run --rm -it --name simpleswap -p 1317:1317/tcp -p 26656:26656/tcp -p 26657:26657/tcp -p 6060:6060/tcp -p 9090:9090/tcp simpleswap:latest 
```


### 3 Executing the use cases

⚠️ replace {servio.address} {virgie.address} {orlando.address} per the address account that you can see in your cli. (step 2)

#### Open a new terminal and put these commands:

##### Providing Liquidity  1/2

```
docker exec -it simpleswap simpleswapd tx simpleswap provide-liquidity 1 1000000000 --from {servio.address}
```

##### Providing Liquidity  2/2

```
docker exec -it simpleswap simpleswapd tx simpleswap provide-liquidity 1 1500000000 --from {virgie.address}
```

#### Seeying the providers liquidity:

```
docker exec -it simpleswap simpleswapd query simpleswap list-pool
```

![image](https://github.com/fenriz07/simpleswap/assets/9199380/85087c2a-6310-4d3e-895b-9f54f4165da0)


#### Checking the back balance
```
docker exec -it simpleswap simpleswapd query bank balances cosmos189z79vlskxjm4n54va5954xlh02ktca6djmct4 
```
![image](https://github.com/fenriz07/simpleswap/assets/9199380/3e12e31e-e493-41bd-9cd4-b2695f3fb966)

#### Swaping eth by usdc

```
docker exec -it simpleswap simpleswapd tx simpleswap swap 100000eth 1  --from {orlando.address}
```

#### Seeying the providers liquidity:
```
docker exec -it simpleswap simpleswapd query simpleswap list-pool
```

![image](https://github.com/fenriz07/simpleswap/assets/9199380/f4c8465f-6d95-4e72-aabb-0ab3e04bb63c)


#### Claiming liquidity and gains by fee

```
docker exec -it simpleswap simpleswapd tx simpleswap claim-liquidity 5 --from {servio.address}
```

#### Checking the servio account balance
```
docker exec -it simpleswap simpleswapd query bank balances {servio.address}
```
![image](https://github.com/fenriz07/simpleswap/assets/9199380/c709e36a-837a-4cff-a78e-abc699d3dc36)


### Seeying the providers liquidity:
```
docker exec -it simpleswap simpleswapd query simpleswap list-pool
```

![image](https://github.com/fenriz07/simpleswap/assets/9199380/f7352462-0f0b-45fc-a2ff-0a8fc3ee2a12)
