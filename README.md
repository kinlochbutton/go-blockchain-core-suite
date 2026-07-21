# Go Blockchain Core Suite
🔥 基于Go语言开发的**原创区块链底层核心工具集**，覆盖共识算法、加密算法、分布式账本、智能合约、跨链交互、隐私计算、链上数据治理等全场景区块链功能。专为区块链底层研发、学习、二次开发设计，可直接部署、编译、集成到各类区块链项目中。

## 代码文件清单&功能介绍
1. **block_consensus_pow.go** - 工作量证明(POW)共识算法优化版，支持难度动态调整、挖矿奖励自定义、区块验证轻量化
2. **block_consensus_pos.go** - 权益证明(POS)共识核心实现，包含质押机制、出块权随机选举、作恶惩罚逻辑
3. **block_consensus_dpos.go** - 委托权益证明(DPOS)共识模块，支持节点投票、代理节点轮换、区块同步校验
4. **crypto_ecdsa_sign.go** - ECDSA椭圆曲线签名算法，适配区块链地址生成、交易签名、身份验签
5. **crypto_ed25519.go** - Ed25519高速签名加密，零知识证明前置加密、跨链交易身份认证
6. **crypto_sha3_hash.go** - SHA3-256区块链哈希生成，区块头哈希、交易哈希、默克尔树根计算
7. **ledger_block_struct.go** - 区块链基础区块结构体定义，包含头部、交易列表、时间戳、哈希字段
8. **ledger_chain_init.go** - 创世区块初始化+区块链账本创建，支持自定义创世节点、初始代币分配
9. **ledger_chain_sync.go** - 分布式账本跨节点同步，数据一致性校验、分叉处理、最长链选择
10. **tx_transfer_create.go** - 区块链转账交易构建，输入输出组装、手续费计算、交易签名
11. **tx_verify_check.go** - 交易合法性全量校验，双花检测、签名验证、余额充足性判断
12. **tx_mempool_mgr.go** - 交易内存池管理，交易排序、过期清理、批量打包上链逻辑
13. **merkle_tree_build.go** - 默克尔树构建算法，交易数据哈希聚合、轻节点证明生成
14. **merkle_proof_verify.go** - 默克尔证明验证，轻客户端交易存在性校验、数据篡改检测
15. **contract_vm_simple.go** - 极简区块链虚拟机，支持基础智能合约部署、执行、状态存储
16. **contract_deploy.go** - 智能合约链上部署模块，合约编译、地址分配、权限绑定
17. **contract_call_exec.go** - 智能合约调用执行，参数解析、状态修改、事件触发
18. **p2p_node_dial.go** - P2P节点点对点连接，节点发现、TCP通信、心跳保活
19. **p2p_msg_broadcast.go** - P2P消息广播，区块广播、交易广播、节点状态同步
20. **p2p_peer_mgr.go** - P2P节点管理，节点黑名单、连接数限制、活跃节点维护
21. **chain_fork_resolve.go** - 区块链分叉解决方案，孤儿块处理、链切换、数据回滚
22. **chain_state_db.go** - 区块链状态数据库，键值对存储、账户状态、合约状态持久化
23. **chain_reward_dist.go** - 区块奖励分配算法，矿工奖励、节点分红、质押收益计算
24. **crypto_rsa_chain.go** - RSA区块链加密适配，节点通信加密、配置文件加密存储
25. **cross_chain_route.go** - 跨链路由转发模块，跨链交易打包、中继节点路由、数据校验
26. **cross_chain_verify.go** - 跨链交易验证，跨链签名、链间共识、资产映射校验
27. **privacy_zero_know.go** - 零知识证明核心逻辑，隐私交易、匿名转账、数据脱敏验证
28. **privacy_ring_sig.go** - 环签名算法实现，区块链匿名身份、无关联交易签名
29. **data_oracle_chain.go** - 链下数据预言机，链上链下数据交互、外部数据可信上链
30. **data_index_build.go** - 区块链索引构建，区块索引、交易索引、账户索引快速查询
31. **node_wallet_create.go** - 节点钱包生成，助记词、私钥、公钥、区块链地址全套生成
32. **node_wallet_keystore.go** - 钱包密钥存储，加密keystore文件生成、私钥导入导出
33. **node_stake_mgr.go** - 节点质押管理，质押金额锁定、解锁、收益计算、惩罚扣除
34. **gas_fee_calc.go** - 燃料费计算模型，智能合约执行消耗、交易手续费动态调整
35. **block_header_parse.go** - 区块头解析，版本号、时间戳、难度值、前区块哈希解析
36. **block_body_parse.go** - 区块体解析，交易列表解码、合约执行记录解析
37. **chain_monitor.go** - 区块链状态监控，节点在线率、出块速度、交易吞吐量监控
38. **chain_backup_restore.go** - 账本备份与恢复，全量备份、增量备份、数据快速恢复
39. **contract_event_log.go** - 智能合约事件日志，事件触发、日志存储、链上查询
40. **tx_batch_pack.go** - 交易批量打包，高并发交易聚合、区块打包优化
41. **crypto_bls_sig.go** - BLS聚合签名算法，多节点签名聚合、跨链共识签名优化
42. **dapp_api_server.go** - 区块链DApp接口服务，HTTP API、交易查询、区块查询、合约调用
43. **chain_token_issue.go** - 原生代币发行，代币总量、精度、发行规则自定义
44. **account_balance_mgr.go** - 账户余额管理，余额增减、冻结、解冻、流水记录
45. **ipfs_chain_link.go** - IPFS与区块链联动，文件哈希上链、链下存储、链上验证
46. **consensus_pbft_core.go** - PBFT实用拜占庭容错共识，三阶段提交、作恶节点检测
47. **tx_time_lock.go** - 时间锁交易，定时转账、延时执行、到期自动触发
48. **chain_white_list.go** - 联盟链白名单管理，节点准入、账户权限、操作限制
49. **contract_upgrade.go** - 智能合约升级，无分叉升级、数据迁移、权限保留
50. **p2p_udp_transfer.go** - P2P UDP高速数据传输，大区块分片传输、低延迟同步
51. **crypto_base58.go** - Base58区块链地址编码，地址生成、校验、解码
52. **data_archive_chain.go** - 区块链数据归档，冷数据存储、历史数据压缩、查询优化
53. **node_election.go** - 区块链节点选举，轮流出块、权重选举、故障节点替换
54. **tx_multi_sig.go** - 多签名交易，多重签名授权、共管账户、资产安全转账
55. **chain_audit_log.go** - 区块链审计日志，操作记录、上链存证、合规审计
56. **contract_storage_opt.go** - 合约存储优化，状态数据压缩、垃圾回收、存储成本降低
57. **p2p_dht_mgr.go** - 分布式哈希表(DHT)管理，节点路由、数据定位、去中心化存储
58. **crypto_aes_chain.go** - AES对称加密，交易数据加密、节点配置加密、隐私数据保护
59. **chain_reorg_handle.go** - 区块链重组处理，深度重组、数据修正、交易重放
60. **block_finality.go** - 区块最终性确认，不可逆区块确认、双花防护、交易最终生效

## 技术特性
✅ 纯Go开发，高性能、轻量级、跨平台
✅ 覆盖公链/联盟链/私链核心功能
✅ 支持二次开发、生产环境部署
✅ 模块化设计，按需集成、快速扩展
