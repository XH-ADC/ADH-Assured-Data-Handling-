		name: "OneToOneCorrespondence",
		args: []string{"fieldA"},
		struc: struct {
			FieldA int `abi:"fieldA"`
  2 changes: 1 addition & 1 deletion2  
accounts/keystore/account_cache_test.go
@@ -318,7 +318,7 @@ func waitForAccounts(wantAccounts []accounts.Account, ks *KeyStore) error {
func TestUpdatedKeyfileContents(t *testing.T) {
	t.Parallel()

	// Create a temporary kesytore to test with
	// Create a temporary keystore to test with
	rand.Seed(time.Now().UnixNano())
	dir := filepath.Join(os.TempDir(), fmt.Sprintf("eth-keystore-updatedkeyfilecontents-test-%d-%d", os.Getpid(), rand.Int()))
	ks := NewKeyStore(dir, LightScryptN, LightScryptP)
  4 changes: 2 additions & 2 deletions4  
accounts/keystore/file_cache.go
@@ -39,7 +39,7 @@ type fileCache struct {
func (fc *fileCache) scan(keyDir string) (mapset.Set, mapset.Set, mapset.Set, error) {
	t0 := time.Now()

	// List all the failes from the keystore folder
	// List all the files from the keystore folder
	files, err := os.ReadDir(keyDir)
	if err != nil {
		return nil, nil, nil, err
@@ -61,7 +61,7 @@ func (fc *fileCache) scan(keyDir string) (mapset.Set, mapset.Set, mapset.Set, er
			log.Trace("Ignoring file on account scan", "path", path)
			continue
		}
		// Gather the set of all and fresly modified files
		// Gather the set of all and freshly modified files
		all.Add(path)

		info, err := fi.Info()
  2 changes: 1 addition & 1 deletion2  
accounts/keystore/keystore_test.go
@@ -214,7 +214,7 @@ func TestSignRace(t *testing.T) {
// Tests that the wallet notifier loop starts and stops correctly based on the
// addition and removal of wallet event subscriptions.
func TestWalletNotifierLifecycle(t *testing.T) {
	// Create a temporary kesytore to test with
	// Create a temporary keystore to test with
	_, ks := tmpKeyStore(t, false)

	// Ensure that the notification updater is not running yet
  4 changes: 2 additions & 2 deletions4  
accounts/usbwallet/trezor.go
@@ -196,10 +196,10 @@ func (w *trezorDriver) trezorDerive(derivationPath []uint32) (common.Address, er
	if _, err := w.trezorExchange(&trezor.EthereumGetAddress{AddressN: derivationPath}, address); err != nil {
		return common.Address{}, err
	}
	if addr := address.GetAddressBin(); len(addr) > 0 { // Older firmwares use binary fomats
	if addr := address.GetAddressBin(); len(addr) > 0 { // Older firmwares use binary formats
		return common.BytesToAddress(addr), nil
	}
	if addr := address.GetAddressHex(); len(addr) > 0 { // Newer firmwares use hexadecimal fomats
	if addr := address.GetAddressHex(); len(addr) > 0 { // Newer firmwares use hexadecimal formats
		return common.HexToAddress(addr), nil
	}
	return common.Address{}, errors.New("missing derived address")
  2 changes: 1 addition & 1 deletion2  
accounts/usbwallet/wallet.go
@@ -380,7 +380,7 @@ func (w *wallet) selfDerive() {
					// of legacy-ledger, the first account on the legacy-path will
					// be shown to the user, even if we don't actively track it
					if i < len(nextAddrs)-1 {
						w.log.Info("Skipping trakcking first account on legacy path, use personal.deriveAccount(<url>,<path>, false) to track",
						w.log.Info("Skipping tracking first account on legacy path, use personal.deriveAccount(<url>,<path>, false) to track",
							"path", path, "address", nextAddrs[i])
						break
					}
  2 changes: 1 addition & 1 deletion2  
build/ci.go
@@ -608,7 +608,7 @@ func doDocker(cmdline []string) {
			}
			if mismatch {
				// Build numbers mismatching, retry in a short time to
				// avoid concurrent failes in both publisher images. If
				// avoid concurrent fails in both publisher images. If
				// however the retry failed too, it means the concurrent
				// builder is still crunching, let that do the publish.
				if i == 0 {
  2 changes: 1 addition & 1 deletion2  
cmd/faucet/faucet.go
@@ -709,7 +709,7 @@ func authTwitter(url string, tokenV1, tokenV2 string) (string, string, string, c
	case tokenV2 != "":
		return authTwitterWithTokenV2(tweetID, tokenV2)
	}
	// Twiter API token isn't provided so we just load the public posts
	// Twitter API token isn't provided so we just load the public posts
	// and scrape it for the Ethereum address and profile URL. We need to load
	// the mobile page though since the main page loads tweet contents via JS.
	url = strings.Replace(url, "https://twitter.com/", "https://mobile.twitter.com/", 1)
  2 changes: 1 addition & 1 deletion2  
cmd/geth/consolecmd_test.go
@@ -155,7 +155,7 @@ To exit, press ctrl-d or type exit
}

// trulyRandInt generates a crypto random integer used by the console tests to
// not clash network ports with other tests running cocurrently.
// not clash network ports with other tests running concurrently.
func trulyRandInt(lo, hi int) int {
	num, _ := rand.Int(rand.Reader, big.NewInt(int64(hi-lo)))
	return int(num.Int64()) + lo
  2 changes: 1 addition & 1 deletion2  
cmd/puppeth/ssh.go
@@ -163,7 +163,7 @@ func dial(server string, pubkey []byte) (*sshClient, error) {
			return nil
		}
		// We have a mismatch, forbid connecting
		return errors.New("ssh key mismatch, readd the machine to update")
		return errors.New("ssh key mismatch, re-add the machine to update")
	}
	client, err := ssh.Dial("tcp", hostport, &ssh.ClientConfig{User: username, Auth: auths, HostKeyCallback: keycheck})
	if err != nil {
  4 changes: 2 additions & 2 deletions4  
common/prque/prque.go
@@ -41,13 +41,13 @@ func (p *Prque) Push(data interface{}, priority int64) {
	heap.Push(p.cont, &item{data, priority})
}

// Peek returns the value with the greates priority but does not pop it off.
// Peek returns the value with the greatest priority but does not pop it off.
func (p *Prque) Peek() (interface{}, int64) {
	item := p.cont.blocks[0][0]
	return item.value, item.priority
}

// Pops the value with the greates priority off the stack and returns it.
// Pops the value with the greatest priority off the stack and returns it.
// Currently no shrinking is done.
func (p *Prque) Pop() (interface{}, int64) {
	item := heap.Pop(p.cont).(*item)
  4 changes: 2 additions & 2 deletions4  
consensus/clique/snapshot_test.go
@@ -305,7 +305,7 @@ func TestClique(t *testing.T) {
		}, {
			// Ensure that pending votes don't survive authorization status changes. This
			// corner case can only appear if a signer is quickly added, removed and then
			// readded (or the inverse), while one of the original voters dropped. If a
			// re-added (or the inverse), while one of the original voters dropped. If a
			// past vote is left cached in the system somewhere, this will interfere with
			// the final signer outcome.
			signers: []string{"A", "B", "C", "D", "E"},
@@ -344,7 +344,7 @@ func TestClique(t *testing.T) {
			},
			failure: errUnauthorizedSigner,
		}, {
			// An authorized signer that signed recenty should not be able to sign again
			// An authorized signer that signed recently should not be able to sign again
			signers: []string{"A", "B"},
			votes: []testerVote{
				{signer: "A"},
  6 changes: 3 additions & 3 deletions6  
console/console.go
@@ -290,7 +290,7 @@ func (c *Console) AutoCompleteInput(line string, pos int) (string, []string, str
	if len(line) == 0 || pos == 0 {
		return "", nil, ""
	}
	// Chunck data to relevant part for autocompletion
	// Chunk data to relevant part for autocompletion
	// E.g. in case of nested lines eth.getBalance(eth.coinb<tab><tab>
	start := pos - 1
	for ; start > 0; start-- {
@@ -407,7 +407,7 @@ func (c *Console) StopInteractive() {
	}
}

// Interactive starts an interactive user session, where in.put is propted from
// Interactive starts an interactive user session, where input is prompted from
// the configured user prompter.
func (c *Console) Interactive() {
	var (
@@ -497,7 +497,7 @@ func (c *Console) readLines(input chan<- string, errc chan<- error, prompt <-cha
	}
}

// countIndents returns the number of identations for the given input.
// countIndents returns the number of indentations for the given input.
// In case of invalid input such as var a = } the result can be negative.
func countIndents(input string) int {
	var (
  4 changes: 2 additions & 2 deletions4  
core/blockchain.go
@@ -1375,7 +1375,7 @@ func (bc *BlockChain) writeBlockAndSetHead(block *types.Block, receipts []*types
		}
		// In theory we should fire a ChainHeadEvent when we inject
		// a canonical block, but sometimes we can insert a batch of
		// canonicial blocks. Avoid firing too many ChainHeadEvents,
		// canonical blocks. Avoid firing too many ChainHeadEvents,
		// we will fire an accumulated ChainHeadEvent and disable fire
		// event here.
		if emitHeadEvent {
@@ -1612,7 +1612,7 @@ func (bc *BlockChain) insertChain(chain types.Blocks, verifySeals, setHead bool)
			// block in the middle. It can only happen in the clique chain. Whenever
			// we insert blocks via `insertSideChain`, we only commit `td`, `header`
			// and `body` if it's non-existent. Since we don't have receipts without
			// reexecution, so nothing to commit. But if the sidechain will be adpoted
			// reexecution, so nothing to commit. But if the sidechain will be adopted
			// as the canonical chain eventually, it needs to be reexecuted for missing
			// state, but if it's this special case here(skip reexecution) we will lose
			// the empty receipt entry.
  32 changes: 16 additions & 16 deletions32  
core/blockchain_repair_test.go
@@ -564,7 +564,7 @@ func testShortReorgedSnapSyncingRepair(t *testing.T, snapshots bool) {
// Tests a recovery for a long canonical chain with frozen blocks where a recent
// block - newer than the ancient limit - was already committed to disk and then
// the process crashed. In this case we expect the chain to be rolled back to the
// committed block, with everything afterwads kept as fast sync data.
// committed block, with everything afterwards kept as fast sync data.
func TestLongShallowRepair(t *testing.T)              { testLongShallowRepair(t, false) }
func TestLongShallowRepairWithSnapshots(t *testing.T) { testLongShallowRepair(t, true) }

@@ -609,7 +609,7 @@ func testLongShallowRepair(t *testing.T, snapshots bool) {
// Tests a recovery for a long canonical chain with frozen blocks where a recent
// block - older than the ancient limit - was already committed to disk and then
// the process crashed. In this case we expect the chain to be rolled back to the
// committed block, with everything afterwads deleted.
// committed block, with everything afterwards deleted.
func TestLongDeepRepair(t *testing.T)              { testLongDeepRepair(t, false) }
func TestLongDeepRepairWithSnapshots(t *testing.T) { testLongDeepRepair(t, true) }

@@ -653,7 +653,7 @@ func testLongDeepRepair(t *testing.T, snapshots bool) {
// Tests a recovery for a long canonical chain with frozen blocks where the fast
// sync pivot point - newer than the ancient limit - was already committed, after
// which the process crashed. In this case we expect the chain to be rolled back
// to the committed block, with everything afterwads kept as fast sync data.
// to the committed block, with everything afterwards kept as fast sync data.
func TestLongSnapSyncedShallowRepair(t *testing.T) {
	testLongSnapSyncedShallowRepair(t, false)
}
@@ -702,7 +702,7 @@ func testLongSnapSyncedShallowRepair(t *testing.T, snapshots bool) {
// Tests a recovery for a long canonical chain with frozen blocks where the fast
// sync pivot point - older than the ancient limit - was already committed, after
// which the process crashed. In this case we expect the chain to be rolled back
// to the committed block, with everything afterwads deleted.
// to the committed block, with everything afterwards deleted.
func TestLongSnapSyncedDeepRepair(t *testing.T)              { testLongSnapSyncedDeepRepair(t, false) }
func TestLongSnapSyncedDeepRepairWithSnapshots(t *testing.T) { testLongSnapSyncedDeepRepair(t, true) }

@@ -843,7 +843,7 @@ func testLongSnapSyncingDeepRepair(t *testing.T, snapshots bool) {
// side chain, where a recent block - newer than the ancient limit - was already
// committed to disk and then the process crashed. In this test scenario the side
// chain is below the committed block. In this case we expect the chain to be
// rolled back to the committed block, with everything afterwads kept as fast
// rolled back to the committed block, with everything afterwards kept as fast
// sync data; the side chain completely nuked by the freezer.
func TestLongOldForkedShallowRepair(t *testing.T) {
	testLongOldForkedShallowRepair(t, false)
@@ -895,7 +895,7 @@ func testLongOldForkedShallowRepair(t *testing.T, snapshots bool) {
// side chain, where a recent block - older than the ancient limit - was already
// committed to disk and then the process crashed. In this test scenario the side
// chain is below the committed block. In this case we expect the canonical chain
// to be rolled back to the committed block, with everything afterwads deleted;
// to be rolled back to the committed block, with everything afterwards deleted;
// the side chain completely nuked by the freezer.
func TestLongOldForkedDeepRepair(t *testing.T)              { testLongOldForkedDeepRepair(t, false) }
func TestLongOldForkedDeepRepairWithSnapshots(t *testing.T) { testLongOldForkedDeepRepair(t, true) }
@@ -942,7 +942,7 @@ func testLongOldForkedDeepRepair(t *testing.T, snapshots bool) {
// side chain, where the fast sync pivot point - newer than the ancient limit -
// was already committed to disk and then the process crashed. In this test scenario
// the side chain is below the committed block. In this case we expect the chain
// to be rolled back to the committed block, with everything afterwads kept as
// to be rolled back to the committed block, with everything afterwards kept as
// fast sync data; the side chain completely nuked by the freezer.
func TestLongOldForkedSnapSyncedShallowRepair(t *testing.T) {
	testLongOldForkedSnapSyncedShallowRepair(t, false)
@@ -994,7 +994,7 @@ func testLongOldForkedSnapSyncedShallowRepair(t *testing.T, snapshots bool) {
// side chain, where the fast sync pivot point - older than the ancient limit -
// was already committed to disk and then the process crashed. In this test scenario
// the side chain is below the committed block. In this case we expect the canonical
// chain to be rolled back to the committed block, with everything afterwads deleted;
// chain to be rolled back to the committed block, with everything afterwards deleted;
// the side chain completely nuked by the freezer.
func TestLongOldForkedSnapSyncedDeepRepair(t *testing.T) {
	testLongOldForkedSnapSyncedDeepRepair(t, false)
@@ -1149,7 +1149,7 @@ func testLongOldForkedSnapSyncingDeepRepair(t *testing.T, snapshots bool) {
// side chain, where a recent block - newer than the ancient limit - was already
// committed to disk and then the process crashed. In this test scenario the side
// chain is above the committed block. In this case we expect the chain to be
// rolled back to the committed block, with everything afterwads kept as fast
// rolled back to the committed block, with everything afterwards kept as fast
// sync data; the side chain completely nuked by the freezer.
func TestLongNewerForkedShallowRepair(t *testing.T) {
	testLongNewerForkedShallowRepair(t, false)
@@ -1201,7 +1201,7 @@ func testLongNewerForkedShallowRepair(t *testing.T, snapshots bool) {
// side chain, where a recent block - older than the ancient limit - was already
// committed to disk and then the process crashed. In this test scenario the side
// chain is above the committed block. In this case we expect the canonical chain
// to be rolled back to the committed block, with everything afterwads deleted;
// to be rolled back to the committed block, with everything afterwards deleted;
// the side chain completely nuked by the freezer.
func TestLongNewerForkedDeepRepair(t *testing.T)              { testLongNewerForkedDeepRepair(t, false) }
func TestLongNewerForkedDeepRepairWithSnapshots(t *testing.T) { testLongNewerForkedDeepRepair(t, true) }
@@ -1248,7 +1248,7 @@ func testLongNewerForkedDeepRepair(t *testing.T, snapshots bool) {
// side chain, where the fast sync pivot point - newer than the ancient limit -
// was already committed to disk and then the process crashed. In this test scenario
// the side chain is above the committed block. In this case we expect the chain
// to be rolled back to the committed block, with everything afterwads kept as fast
// to be rolled back to the committed block, with everything afterwards kept as fast
// sync data; the side chain completely nuked by the freezer.
func TestLongNewerForkedSnapSyncedShallowRepair(t *testing.T) {
	testLongNewerForkedSnapSyncedShallowRepair(t, false)
@@ -1300,7 +1300,7 @@ func testLongNewerForkedSnapSyncedShallowRepair(t *testing.T, snapshots bool) {
// side chain, where the fast sync pivot point - older than the ancient limit -
// was already committed to disk and then the process crashed. In this test scenario
// the side chain is above the committed block. In this case we expect the canonical
// chain to be rolled back to the committed block, with everything afterwads deleted;
// chain to be rolled back to the committed block, with everything afterwards deleted;
// the side chain completely nuked by the freezer.
func TestLongNewerForkedSnapSyncedDeepRepair(t *testing.T) {
	testLongNewerForkedSnapSyncedDeepRepair(t, false)
@@ -1454,7 +1454,7 @@ func testLongNewerForkedSnapSyncingDeepRepair(t *testing.T, snapshots bool) {
// Tests a recovery for a long canonical chain with frozen blocks and a longer side
// chain, where a recent block - newer than the ancient limit - was already committed
// to disk and then the process crashed. In this case we expect the chain to be
// rolled back to the committed block, with everything afterwads kept as fast sync
// rolled back to the committed block, with everything afterwards kept as fast sync
// data. The side chain completely nuked by the freezer.
func TestLongReorgedShallowRepair(t *testing.T)              { testLongReorgedShallowRepair(t, false) }
func TestLongReorgedShallowRepairWithSnapshots(t *testing.T) { testLongReorgedShallowRepair(t, true) }
@@ -1501,7 +1501,7 @@ func testLongReorgedShallowRepair(t *testing.T, snapshots bool) {
// Tests a recovery for a long canonical chain with frozen blocks and a longer side
// chain, where a recent block - older than the ancient limit - was already committed
// to disk and then the process crashed. In this case we expect the canonical chains
// to be rolled back to the committed block, with everything afterwads deleted. The
// to be rolled back to the committed block, with everything afterwards deleted. The
// side chain completely nuked by the freezer.
func TestLongReorgedDeepRepair(t *testing.T)              { testLongReorgedDeepRepair(t, false) }
func TestLongReorgedDeepRepairWithSnapshots(t *testing.T) { testLongReorgedDeepRepair(t, true) }
@@ -1548,7 +1548,7 @@ func testLongReorgedDeepRepair(t *testing.T, snapshots bool) {
// side chain, where the fast sync pivot point - newer than the ancient limit -
// was already committed to disk and then the process crashed. In this case we
// expect the chain to be rolled back to the committed block, with everything
// afterwads kept as fast sync data. The side chain completely nuked by the
// afterwards kept as fast sync data. The side chain completely nuked by the
// freezer.
func TestLongReorgedSnapSyncedShallowRepair(t *testing.T) {
	testLongReorgedSnapSyncedShallowRepair(t, false)
@@ -1600,7 +1600,7 @@ func testLongReorgedSnapSyncedShallowRepair(t *testing.T, snapshots bool) {
// side chain, where the fast sync pivot point - older than the ancient limit -
// was already committed to disk and then the process crashed. In this case we
// expect the canonical chains to be rolled back to the committed block, with
// everything afterwads deleted. The side chain completely nuked by the freezer.
// everything afterwards deleted. The side chain completely nuked by the freezer.
func TestLongReorgedSnapSyncedDeepRepair(t *testing.T) {
	testLongReorgedSnapSyncedDeepRepair(t, false)
}
  2 changes: 1 addition & 1 deletion2  
core/rawdb/accessors_chain_test.go
@@ -285,7 +285,7 @@ func TestTdStorage(t *testing.T) {
func TestCanonicalMappingStorage(t *testing.T) {
	db := NewMemoryDatabase()

	// Create a test canonical number and assinged hash to move around
	// Create a test canonical number and assigned hash to move around
	hash, number := common.Hash{0: 0xff}, uint64(314)
	if entry := ReadCanonicalHash(db, number); entry != (common.Hash{}) {
		t.Fatalf("Non existent canonical mapping returned: %v", entry)
  2 changes: 1 addition & 1 deletion2  
core/rawdb/database.go
@@ -260,7 +260,7 @@ func NewDatabaseWithFreezer(db ethdb.KeyValueStore, ancient string, namespace st
				if kvblob, _ := db.Get(headerHashKey(1)); len(kvblob) == 0 {
					return nil, errors.New("ancient chain segments already extracted, please set --datadir.ancient to the correct path")
				}
				// Block #1 is still in the database, we're allowed to init a new feezer
				// Block #1 is still in the database, we're allowed to init a new freezer
			}
			// Otherwise, the head header is still the genesis, we're allowed to init a new
			// freezer.
  2 changes: 1 addition & 1 deletion2  
core/rawdb/freezer_table.go
@@ -46,7 +46,7 @@ var (
	errNotSupported = errors.New("this operation is not supported")
)

// indexEntry contains the number/id of the file that the data resides in, aswell as the
// indexEntry contains the number/id of the file that the data resides in, as well as the
// offset within the file to the end of the data.
// In serialized form, the filenum is stored as uint16.
type indexEntry struct {
  4 changes: 2 additions & 2 deletions4  
core/state/snapshot/iterator_fast.go
@@ -319,15 +319,15 @@ func (fi *fastIterator) Slot() []byte {
}

// Release iterates over all the remaining live layer iterators and releases each
// of thme individually.
// of them individually.
func (fi *fastIterator) Release() {
	for _, it := range fi.iterators {
		it.it.Release()
	}
	fi.iterators = nil
}

// Debug is a convencience helper during testing
// Debug is a convenience helper during testing
func (fi *fastIterator) Debug() {
	for _, it := range fi.iterators {
		fmt.Printf("[p=%v v=%v] ", it.priority, it.it.Hash()[0])
  2 changes: 1 addition & 1 deletion2  
core/state/snapshot/snapshot_test.go
@@ -265,7 +265,7 @@ func TestPostCapBasicDataAccess(t *testing.T) {
	snaps.Update(common.HexToHash("0xa3"), common.HexToHash("0xa2"), nil, setAccount("0xa3"), nil)
	snaps.Update(common.HexToHash("0xb3"), common.HexToHash("0xb2"), nil, setAccount("0xb3"), nil)

	// checkExist verifies if an account exiss in a snapshot
	// checkExist verifies if an account exists in a snapshot
	checkExist := func(layer *diffLayer, key string) error {
		if data, _ := layer.Account(common.HexToHash(key)); data == nil {
			return fmt.Errorf("expected %x to exist, got nil", common.HexToHash(key))
  6 changes: 3 additions & 3 deletions6  
core/state/statedb.go
@@ -792,7 +792,7 @@ func (s *StateDB) Finalise(deleteEmptyObjects bool) {
			// If state snapshotting is active, also mark the destruction there.
			// Note, we can't do this only at the end of a block because multiple
			// transactions within the same block might self destruct and then
			// ressurrect an account; but the snapshotter needs both events.
			// resurrect an account; but the snapshotter needs both events.
			if s.snap != nil {
				s.snapDestructs[obj.addrHash] = struct{}{} // We need to maintain account deletions explicitly (will remain set indefinitely)
				delete(s.snapAccounts, obj.addrHash)       // Clear out any previously updated account data (may be recreated via a ressurrect)
@@ -891,7 +891,7 @@ func (s *StateDB) clearJournalAndRefund() {
		s.journal = newJournal()
		s.refund = 0
	}
	s.validRevisions = s.validRevisions[:0] // Snapshots can be created without journal entires
	s.validRevisions = s.validRevisions[:0] // Snapshots can be created without journal entries
}

// Commit writes the state to the underlying in-memory trie database.
@@ -938,7 +938,7 @@ func (s *StateDB) Commit(deleteEmptyObjects bool) (common.Hash, error) {
			log.Crit("Failed to commit dirty codes", "error", err)
		}
	}
	// Write the account trie changes, measuing the amount of wasted time
	// Write the account trie changes, measuring the amount of wasted time
	var start time.Time
	if metrics.EnabledExpensive {
		start = time.Now()
  2 changes: 1 addition & 1 deletion2  
core/state/statedb_test.go
@@ -771,7 +771,7 @@ func TestStateDBAccessList(t *testing.T) {
				t.Fatalf("expected %x to be in access list", address)
			}
		}
		// Check that only the expected addresses are present in the acesslist
		// Check that only the expected addresses are present in the access list
		for address := range state.accessList.addresses {
			if _, exist := addressMap[address]; !exist {
				t.Fatalf("extra address %x in access list", address)
  12 changes: 6 additions & 6 deletions12  
core/state/sync_test.go
@@ -305,8 +305,8 @@ func TestIterativeDelayedStateSync(t *testing.T) {
	}
	for len(nodeElements)+len(codeElements) > 0 {
		// Sync only half of the scheduled nodes
		var nodeProcessd int
		var codeProcessd int
		var nodeProcessed int
		var codeProcessed int
		if len(codeElements) > 0 {
			codeResults := make([]trie.CodeSyncResult, len(codeElements)/2+1)
			for i, element := range codeElements[:len(codeResults)] {
@@ -321,7 +321,7 @@ func TestIterativeDelayedStateSync(t *testing.T) {
					t.Fatalf("failed to process result %v", err)
				}
			}
			codeProcessd = len(codeResults)
			codeProcessed = len(codeResults)
		}
		if len(nodeElements) > 0 {
			nodeResults := make([]trie.NodeSyncResult, len(nodeElements)/2+1)
@@ -337,7 +337,7 @@ func TestIterativeDelayedStateSync(t *testing.T) {
					t.Fatalf("failed to process result %v", err)
				}
			}
			nodeProcessd = len(nodeResults)
			nodeProcessed = len(nodeResults)
		}
		batch := dstDb.NewBatch()
		if err := sched.Commit(batch); err != nil {
@@ -346,15 +346,15 @@ func TestIterativeDelayedStateSync(t *testing.T) {
		batch.Write()

		paths, nodes, codes = sched.Missing(0)
		nodeElements = nodeElements[nodeProcessd:]
		nodeElements = nodeElements[nodeProcessed:]
		for i := 0; i < len(paths); i++ {
			nodeElements = append(nodeElements, stateElement{
				path:     paths[i],
				hash:     nodes[i],
				syncPath: trie.NewSyncPath([]byte(paths[i])),
			})
		}
		codeElements = codeElements[codeProcessd:]
		codeElements = codeElements[codeProcessed:]
		for i := 0; i < len(codes); i++ {
			codeElements = append(codeElements, stateElement{
				code: codes[i],
  2 changes: 1 addition & 1 deletion2  
core/state/trie_prefetcher.go
@@ -212,7 +212,7 @@ type subfetcher struct {

	wake chan struct{}  // Wake channel if a new task is scheduled
	stop chan struct{}  // Channel to interrupt processing
	term chan struct{}  // Channel to signal iterruption
	term chan struct{}  // Channel to signal interruption
	copy chan chan Trie // Channel to request a copy of the current trie

	seen map[string]struct{} // Tracks the entries already loaded
  2 changes: 1 addition & 1 deletion2  
core/types/block.go
@@ -317,7 +317,7 @@ func (b *Block) Header() *Header { return CopyHeader(b.header) }
func (b *Block) Body() *Body { return &Body{b.transactions, b.uncles} }

// Size returns the true RLP encoded storage size of the block, either by encoding
// and returning it, or returning a previsouly cached value.
// and returning it, or returning a previously cached value.
func (b *Block) Size() common.StorageSize {
	if size := b.size.Load(); size != nil {
		return size.(common.StorageSize)
  2 changes: 1 addition & 1 deletion2  
core/types/block_test.go
@@ -314,7 +314,7 @@ func TestRlpDecodeParentHash(t *testing.T) {
	}
	// Also test a very very large header.
	{
		// The rlp-encoding of the heder belowCauses _total_ length of 65540,
		// The rlp-encoding of the header belowCauses _total_ length of 65540,
		// which is the first to blow the fast-path.
		h := &Header{
			ParentHash: want,
  4 changes: 2 additions & 2 deletions4  
core/vm/runtime/runtime_test.go
@@ -457,7 +457,7 @@ func BenchmarkSimpleLoop(b *testing.B) {
		byte(vm.JUMP),
	}

	calllRevertingContractWithInput := []byte{
	callRevertingContractWithInput := []byte{
		byte(vm.JUMPDEST), //
		// push args for the call
		byte(vm.PUSH1), 0, // out size
@@ -485,7 +485,7 @@ func BenchmarkSimpleLoop(b *testing.B) {
	benchmarkNonModifyingCode(100000000, loopingCode, "loop-100M", "", b)
	benchmarkNonModifyingCode(100000000, callInexistant, "call-nonexist-100M", "", b)
	benchmarkNonModifyingCode(100000000, callEOA, "call-EOA-100M", "", b)
	benchmarkNonModifyingCode(100000000, calllRevertingContractWithInput, "call-reverting-100M", "", b)
	benchmarkNonModifyingCode(100000000, callRevertingContractWithInput, "call-reverting-100M", "", b)

	//benchmarkNonModifyingCode(10000000, staticCallIdentity, "staticcall-identity-10M", b)
	//benchmarkNonModifyingCode(10000000, loopingCode, "loop-10M", b)
  4 changes: 2 additions & 2 deletions4  
crypto/bls12381/isogeny.go
@@ -19,7 +19,7 @@ package bls12381
// isogenyMapG1 applies 11-isogeny map for BLS12-381 G1 defined at draft-irtf-cfrg-hash-to-curve-06.
func isogenyMapG1(x, y *fe) {
	// https://tools.ietf.org/html/draft-irtf-cfrg-hash-to-curve-06#appendix-C.2
	params := isogenyConstansG1
	params := isogenyConstantsG1
	degree := 15
	xNum, xDen, yNum, yDen := new(fe), new(fe), new(fe), new(fe)
	xNum.set(params[0][degree])
@@ -76,7 +76,7 @@ func isogenyMapG2(e *fp2, x, y *fe2) {
	y.set(yNum)
}

var isogenyConstansG1 = [4][16]*fe{
var isogenyConstantsG1 = [4][16]*fe{
	{
		{0x4d18b6f3af00131c, 0x19fa219793fee28c, 0x3f2885f1467f19ae, 0x23dcea34f2ffb304, 0xd15b58d2ffc00054, 0x0913be200a20bef4},
		{0x898985385cdbbd8b, 0x3c79e43cc7d966aa, 0x1597e193f4cd233a, 0x8637ef1e4d6623ad, 0x11b22deed20d827b, 0x07097bc5998784ad},
  4 changes: 2 additions & 2 deletions4  
eth/catalyst/api.go
@@ -539,7 +539,7 @@ func (api *ConsensusAPI) invalid(err error, latestValid *types.Header) beacon.Pa
	return beacon.PayloadStatusV1{Status: beacon.INVALID, LatestValidHash: &currentHash, ValidationError: &errorMsg}
}

// heatbeat loops indefinitely, and checks if there have been beacon client updates
// heartbeat loops indefinitely, and checks if there have been beacon client updates
// received in the last while. If not - or if they but strange ones - it warns the
// user that something might be off with their consensus node.
//
@@ -649,7 +649,7 @@ func (api *ConsensusAPI) heartbeat() {
					if eta == 0 {
						log.Warn(message)
					} else {
						log.Warn(message, "eta", common.PrettyAge(time.Now().Add(-eta))) // weird hack, but duration formatted doens't handle days
						log.Warn(message, "eta", common.PrettyAge(time.Now().Add(-eta))) // weird hack, but duration formatted doesn't handle days
					}
					offlineLogged = time.Now()
				}
  2 changes: 1 addition & 1 deletion2  
eth/downloader/api.go
@@ -125,7 +125,7 @@ type SyncingResult struct {
	Status  ethereum.SyncProgress `json:"status"`
}

// uninstallSyncSubscriptionRequest uninstalles a syncing subscription in the API event loop.
// uninstallSyncSubscriptionRequest uninstalls a syncing subscription in the API event loop.
type uninstallSyncSubscriptionRequest struct {
	c           chan interface{}
	uninstalled chan interface{}
  2 changes: 1 addition & 1 deletion2  
eth/downloader/beaconsync.go
@@ -236,7 +236,7 @@ func (d *Downloader) findBeaconAncestor() (uint64, error) {
	// Binary search to find the ancestor
	start, end := beaconTail.Number.Uint64()-1, number
	if number := beaconHead.Number.Uint64(); end > number {
		// This shouldn't really happen in a healty network, but if the consensus
		// This shouldn't really happen in a healthy network, but if the consensus
		// clients feeds us a shorter chain as the canonical, we should not attempt
		// to access non-existent skeleton items.
		log.Warn("Beacon head lower than local chain", "beacon", number, "local", end)
  4 changes: 2 additions & 2 deletions4  
eth/downloader/downloader.go
@@ -364,7 +364,7 @@ func (d *Downloader) synchronise(id string, hash common.Hash, td, ttd *big.Int,
	// The beacon header syncer is async. It will start this synchronization and
	// will continue doing other tasks. However, if synchronization needs to be
	// cancelled, the syncer needs to know if we reached the startup point (and
	// inited the cancel cannel) or not yet. Make sure that we'll signal even in
	// inited the cancel channel) or not yet. Make sure that we'll signal even in
	// case of a failure.
	if beaconPing != nil {
		defer func() {
@@ -1461,7 +1461,7 @@ func (d *Downloader) processHeaders(origin uint64, td, ttd *big.Int, beaconMode
			}
			d.syncStatsLock.Unlock()

			// Signal the content downloaders of the availablility of new tasks
			// Signal the content downloaders of the availability of new tasks
			for _, ch := range []chan bool{d.queue.blockWakeCh, d.queue.receiptWakeCh} {
				select {
				case ch <- true:
  8 changes: 4 additions & 4 deletions8  
eth/downloader/downloader_test.go
@@ -360,7 +360,7 @@ func (dlp *downloadTesterPeer) RequestAccountRange(id uint64, root, origin, limi
}

// RequestStorageRanges fetches a batch of storage slots belonging to one or
// more accounts. If slots from only one accout is requested, an origin marker
// more accounts. If slots from only one account is requested, an origin marker
// may also be used to retrieve from there.
func (dlp *downloadTesterPeer) RequestStorageRanges(id uint64, root common.Hash, accounts []common.Hash, origin, limit []byte, bytes uint64) error {
	// Create the request and service it
@@ -399,7 +399,7 @@ func (dlp *downloadTesterPeer) RequestByteCodes(id uint64, hashes []common.Hash,
}

// RequestTrieNodes fetches a batch of account or storage trie nodes rooted in
// a specificstate trie.
// a specific state trie.
func (dlp *downloadTesterPeer) RequestTrieNodes(id uint64, root common.Hash, paths []snap.TrieNodePathSet, bytes uint64) error {
	req := &snap.GetTrieNodesPacket{
		ID:    id,
@@ -571,8 +571,8 @@ func testForkedSync(t *testing.T, protocol uint, mode SyncMode) {
	assertOwnChain(t, tester, len(chainB.blocks))
}

// Tests that synchronising against a much shorter but much heavyer fork works
// corrently and is not dropped.
// Tests that synchronising against a much shorter but much heavier fork works
// currently and is not dropped.
func TestHeavyForkedSync66Full(t *testing.T)  { testHeavyForkedSync(t, eth.ETH66, FullSync) }
func TestHeavyForkedSync66Snap(t *testing.T)  { testHeavyForkedSync(t, eth.ETH66, SnapSync) }
func TestHeavyForkedSync66Light(t *testing.T) { testHeavyForkedSync(t, eth.ETH66, LightSync) }
  6 changes: 3 additions & 3 deletions6  
eth/downloader/fetchers_concurrent.go
@@ -47,7 +47,7 @@ type typedQueue interface {

	// capacity is responsible for calculating how many items of the abstracted
	// type a particular peer is estimated to be able to retrieve within the
	// alloted round trip time.
	// allotted round trip time.
	capacity(peer *peerConnection, rtt time.Duration) int

	// updateCapacity is responsible for updating how many items of the abstracted
@@ -58,7 +58,7 @@ type typedQueue interface {
	// from the download queue to the specified peer.
	reserve(peer *peerConnection, items int) (*fetchRequest, bool, bool)

	// unreserve is resposible for removing the current retrieval allocation
	// unreserve is responsible for removing the current retrieval allocation
	// assigned to a specific peer and placing it back into the pool to allow
	// reassigning to some other peer.
	unreserve(peer string) int
@@ -190,7 +190,7 @@ func (d *Downloader) concurrentFetch(queue typedQueue, beaconMode bool) error {
				req, err := queue.request(peer, request, responses)
				if err != nil {
					// Sending the request failed, which generally means the peer
					// was diconnected in between assignment and network send.
					// was disconnected in between assignment and network send.
					// Although all peer removal operations return allocated tasks
					// to the queue, that is async, and we can do better here by
					// immediately pushing the unfulfilled requests.
  4 changes: 2 additions & 2 deletions4  
eth/downloader/fetchers_concurrent_bodies.go
@@ -41,7 +41,7 @@ func (q *bodyQueue) pending() int {
}

// capacity is responsible for calculating how many bodies a particular peer is
// estimated to be able to retrieve within the alloted round trip time.
// estimated to be able to retrieve within the allotted round trip time.
func (q *bodyQueue) capacity(peer *peerConnection, rtt time.Duration) int {
	return peer.BodyCapacity(rtt)
}
@@ -58,7 +58,7 @@ func (q *bodyQueue) reserve(peer *peerConnection, items int) (*fetchRequest, boo
	return q.queue.ReserveBodies(peer, items)
}

// unreserve is resposible for removing the current body retrieval allocation
// unreserve is responsible for removing the current body retrieval allocation
// assigned to a specific peer and placing it back into the pool to allow
// reassigning to some other peer.
func (q *bodyQueue) unreserve(peer string) int {
  4 changes: 2 additions & 2 deletions4  
eth/downloader/fetchers_concurrent_headers.go
@@ -41,7 +41,7 @@ func (q *headerQueue) pending() int {
}

// capacity is responsible for calculating how many headers a particular peer is
// estimated to be able to retrieve within the alloted round trip time.
// estimated to be able to retrieve within the allotted round trip time.
func (q *headerQueue) capacity(peer *peerConnection, rtt time.Duration) int {
	return peer.HeaderCapacity(rtt)
}
@@ -58,7 +58,7 @@ func (q *headerQueue) reserve(peer *peerConnection, items int) (*fetchRequest, b
	return q.queue.ReserveHeaders(peer, items), false, false
}

// unreserve is resposible for removing the current header retrieval allocation
// unreserve is responsible for removing the current header retrieval allocation
// assigned to a specific peer and placing it back into the pool to allow
// reassigning to some other peer.
func (q *headerQueue) unreserve(peer string) int {
  6 changes: 3 additions & 3 deletions6  
eth/downloader/fetchers_concurrent_receipts.go
@@ -28,7 +28,7 @@ import (
// concurrent fetcher and the downloader.
type receiptQueue Downloader

// waker returns a notification channel that gets pinged in case more reecipt
// waker returns a notification channel that gets pinged in case more receipt
// fetches have been queued up, so the fetcher might assign it to idle peers.
func (q *receiptQueue) waker() chan bool {
	return q.queue.receiptWakeCh
@@ -41,7 +41,7 @@ func (q *receiptQueue) pending() int {
}

// capacity is responsible for calculating how many receipts a particular peer is
// estimated to be able to retrieve within the alloted round trip time.
// estimated to be able to retrieve within the allotted round trip time.
func (q *receiptQueue) capacity(peer *peerConnection, rtt time.Duration) int {
	return peer.ReceiptCapacity(rtt)
}
@@ -58,7 +58,7 @@ func (q *receiptQueue) reserve(peer *peerConnection, items int) (*fetchRequest,
	return q.queue.ReserveReceipts(peer, items)
}

// unreserve is resposible for removing the current receipt retrieval allocation
// unreserve is responsible for removing the current receipt retrieval allocation
// assigned to a specific peer and placing it back into the pool to allow
// reassigning to some other peer.
func (q *receiptQueue) unreserve(peer string) int {
  2 changes: 1 addition & 1 deletion2  
eth/downloader/queue.go
@@ -859,7 +859,7 @@ func (q *queue) deliver(id string, taskPool map[common.Hash]*types.Header,
		if res, stale, err := q.resultCache.GetDeliverySlot(header.Number.Uint64()); err == nil {
			reconstruct(accepted, res)
		} else {
			// else: betweeen here and above, some other peer filled this result,
			// else: between here and above, some other peer filled this result,
			// or it was indeed a no-op. This should not happen, but if it does it's
			// not something to panic about
			log.Error("Delivery stale", "stale", stale, "number", header.Number.Uint64(), "err", err)
  12 changes: 6 additions & 6 deletions12  
eth/downloader/skeleton.go
@@ -51,7 +51,7 @@ const requestHeaders = 512
// errSyncLinked is an internal helper error to signal that the current sync
// cycle linked up to the genesis block, this the skeleton syncer should ping
// the backfiller to resume. Since we already have that logic on sync start,
// piggie-back on that instead of 2 entrypoints.
// piggy-back on that instead of 2 entrypoints.
var errSyncLinked = errors.New("sync linked")

// errSyncMerged is an internal helper error to signal that the current sync
@@ -148,7 +148,7 @@ type backfiller interface {
	// suspend requests the backfiller to abort any running full or snap sync
	// based on the skeleton chain as it might be invalid. The backfiller should
	// gracefully handle multiple consecutive suspends without a resume, even
	// on initial sartup.
	// on initial startup.
	//
	// The method should return the last block header that has been successfully
	// backfilled, or nil if the backfiller was not resumed.
@@ -209,7 +209,7 @@ type skeleton struct {

	headEvents chan *headUpdate // Notification channel for new heads
	terminate  chan chan error  // Termination channel to abort sync
	terminated chan struct{}    // Channel to signal that the syner is dead
	terminated chan struct{}    // Channel to signal that the syncer is dead

	// Callback hooks used during testing
	syncStarting func() // callback triggered after a sync cycle is inited but before started
@@ -553,7 +553,7 @@ func (s *skeleton) initSync(head *types.Header) {
			return
		}
	}
	// Either we've failed to decode the previus state, or there was none. Start
	// Either we've failed to decode the previous state, or there was none. Start
	// a fresh sync with a single subchain represented by the currently sent
	// chain head.
	s.progress = &skeletonProgress{
@@ -823,7 +823,7 @@ func (s *skeleton) executeTask(peer *peerConnection, req *headerRequest) {
	}
}

// revertRequests locates all the currently pending reuqests from a particular
// revertRequests locates all the currently pending requests from a particular
// peer and reverts them, rescheduling for others to fulfill.
func (s *skeleton) revertRequests(peer string) {
	// Gather the requests first, revertals need the lock too
@@ -871,7 +871,7 @@ func (s *skeleton) revertRequest(req *headerRequest) {
	delete(s.requests, req.id)

	// Remove the request from the tracked set and mark the task as not-pending,
	// ready for resheduling
	// ready for rescheduling
	s.scratchOwners[(s.scratchHead-req.head)/requestHeaders] = ""
}

  14 changes: 7 additions & 7 deletions14  
eth/downloader/skeleton_test.go
@@ -53,7 +53,7 @@ func newHookedBackfiller() backfiller {
// suspend requests the backfiller to abort any running full or snap sync
// based on the skeleton chain as it might be invalid. The backfiller should
// gracefully handle multiple consecutive suspends without a resume, even
// on initial sartup.
// on initial startup.
func (hf *hookedBackfiller) suspend() *types.Header {
	if hf.suspendHook != nil {
		hf.suspendHook()
@@ -111,7 +111,7 @@ func newSkeletonTestPeerWithHook(id string, headers []*types.Header, serve func(
// function can be used to retrieve batches of headers from the particular peer.
func (p *skeletonTestPeer) RequestHeadersByNumber(origin uint64, amount int, skip int, reverse bool, sink chan *eth.Response) (*eth.Request, error) {
	// Since skeleton test peer are in-memory mocks, dropping the does not make
	// them inaccepssible. As such, check a local `dropped` field to see if the
	// them inaccessible. As such, check a local `dropped` field to see if the
	// peer has been dropped and should not respond any more.
	if atomic.LoadUint64(&p.dropped) != 0 {
		return nil, errors.New("peer already dropped")
@@ -204,7 +204,7 @@ func (p *skeletonTestPeer) RequestReceipts([]common.Hash, chan *eth.Response) (*
	panic("skeleton sync must not request receipts")
}

// Tests various sync initialzations based on previous leftovers in the database
// Tests various sync initializations based on previous leftovers in the database
// and announced heads.
func TestSkeletonSyncInit(t *testing.T) {
	// Create a few key headers
@@ -227,7 +227,7 @@ func TestSkeletonSyncInit(t *testing.T) {
			newstate: []*subchain{{Head: 50, Tail: 50}},
		},
		// Empty database with only the genesis set with a leftover empty sync
		// progess. This is a synthetic case, just for the sake of covering things.
		// progress. This is a synthetic case, just for the sake of covering things.
		{
			oldstate: []*subchain{},
			head:     block50,
@@ -533,13 +533,13 @@ func TestSkeletonSyncRetrievals(t *testing.T) {
		peers    []*skeletonTestPeer // Initial peer set to start the sync with
		midstate []*subchain         // Expected sync state after initial cycle
		midserve uint64              // Expected number of header retrievals after initial cycle
		middrop  uint64              // Expectd number of peers dropped after initial cycle
		middrop  uint64              // Expected number of peers dropped after initial cycle

		newHead  *types.Header     // New header to annount on top of the old one
		newHead  *types.Header     // New header to anoint on top of the old one
		newPeer  *skeletonTestPeer // New peer to join the skeleton syncer
		endstate []*subchain       // Expected sync state after the post-init event
		endserve uint64            // Expected number of header retrievals after the post-init event
		enddrop  uint64            // Expectd number of peers dropped after the post-init event
		enddrop  uint64            // Expected number of peers dropped after the post-init event
	}{
		// Completely empty database with only the genesis set. The sync is expected
		// to create a single subchain with the requested head. No peers however, so
  2 changes: 1 addition & 1 deletion2  
eth/ethconfig/config.go
@@ -196,7 +196,7 @@ type Config struct {
	RPCEVMTimeout time.Duration

	// RPCTxFeeCap is the global transaction fee(price * gaslimit) cap for
	// send-transction variants. The unit is ether.
	// send-transaction variants. The unit is ether.
	RPCTxFeeCap float64

	// Checkpoint is a hardcoded checkpoint which can be nil.
  10 changes: 5 additions & 5 deletions10  
eth/fetcher/tx_fetcher.go
@@ -120,7 +120,7 @@ type txDelivery struct {
	direct bool          // Whether this is a direct reply or a broadcast
}

// txDrop is the notiication that a peer has disconnected.
// txDrop is the notification that a peer has disconnected.
type txDrop struct {
	peer string
}
@@ -260,7 +260,7 @@ func (f *TxFetcher) Notify(peer string, hashes []common.Hash) error {
// Enqueue imports a batch of received transaction into the transaction pool
// and the fetcher. This method may be called by both transaction broadcasts and
// direct request replies. The differentiation is important so the fetcher can
// re-shedule missing transactions as soon as possible.
// re-schedule missing transactions as soon as possible.
func (f *TxFetcher) Enqueue(peer string, txs []*types.Transaction, direct bool) error {
	// Keep track of all the propagated transactions
	if direct {
@@ -558,7 +558,7 @@ func (f *TxFetcher) loop() {
			// In case of a direct delivery, also reschedule anything missing
			// from the original query
			if delivery.direct {
				// Mark the reqesting successful (independent of individual status)
				// Mark the requesting successful (independent of individual status)
				txRequestDoneMeter.Mark(int64(len(delivery.hashes)))

				// Make sure something was pending, nuke it
@@ -607,7 +607,7 @@ func (f *TxFetcher) loop() {
					delete(f.alternates, hash)
					delete(f.fetching, hash)
				}
				// Something was delivered, try to rechedule requests
				// Something was delivered, try to reschedule requests
				f.scheduleFetches(timeoutTimer, timeoutTrigger, nil) // Partial delivery may enable others to deliver too
			}

@@ -719,7 +719,7 @@ func (f *TxFetcher) rescheduleWait(timer *mclock.Timer, trigger chan struct{}) {
// should be rescheduled if some request is pending. In practice, a timeout will
// cause the timer to be rescheduled every 5 secs (until the peer comes through or
// disconnects). This is a limitation of the fetcher code because we don't trac
// pending requests and timed out requests separatey. Without double tracking, if
// pending requests and timed out requests separately. Without double tracking, if
// we simply didn't reschedule the timer on all-timeout then the timer would never
// be set again since len(request) > 0 => something's running.
func (f *TxFetcher) rescheduleTimeout(timer *mclock.Timer, trigger chan struct{}) {
  6 changes: 3 additions & 3 deletions6  
eth/fetcher/tx_fetcher_test.go
@@ -1011,7 +1011,7 @@ func TestTransactionFetcherOutOfBoundDeliveries(t *testing.T) {
}

// Tests that dropping a peer cleans out all internal data structures in all the
// live or danglng stages.
// live or dangling stages.
func TestTransactionFetcherDrop(t *testing.T) {
	testTransactionFetcherParallel(t, txFetcherTest{
		init: func() *TxFetcher {
@@ -1121,7 +1121,7 @@ func TestTransactionFetcherDropRescheduling(t *testing.T) {
}

// This test reproduces a crash caught by the fuzzer. The root cause was a
// dangling transaction timing out and clashing on readd with a concurrently
// dangling transaction timing out and clashing on re-add with a concurrently
// announced one.
func TestTransactionFetcherFuzzCrash01(t *testing.T) {
	testTransactionFetcherParallel(t, txFetcherTest{
@@ -1148,7 +1148,7 @@ func TestTransactionFetcherFuzzCrash01(t *testing.T) {
}

// This test reproduces a crash caught by the fuzzer. The root cause was a
// dangling transaction getting peer-dropped and clashing on readd with a
// dangling transaction getting peer-dropped and clashing on re-add with a
// concurrently announced one.
func TestTransactionFetcherFuzzCrash02(t *testing.T) {
	testTransactionFetcherParallel(t, txFetcherTest{
  2 changes: 1 addition & 1 deletion2  
eth/filters/api.go
@@ -36,7 +36,7 @@ import (
// and associated subscription in the event system.
type filter struct {
	typ      Type
	deadline *time.Timer // filter is inactiv when deadline triggers
	deadline *time.Timer // filter is inactive when deadline triggers
	hashes   []common.Hash
	crit     FilterCriteria
	logs     []*types.Log
  2 changes: 1 addition & 1 deletion2  
eth/peerset.go
@@ -41,7 +41,7 @@ var (
	errPeerNotRegistered = errors.New("peer not registered")

	// errSnapWithoutEth is returned if a peer attempts to connect only on the
	// snap protocol without advertizing the eth main protocol.
	// snap protocol without advertising the eth main protocol.
	errSnapWithoutEth = errors.New("peer connected on snap without compatible eth support")
)

  2 changes: 1 addition & 1 deletion2  
eth/protocols/eth/broadcast.go
@@ -36,7 +36,7 @@ type blockPropagation struct {
	td    *big.Int
}

// broadcastBlocks is a write loop that multiplexes blocks and block accouncements
// broadcastBlocks is a write loop that multiplexes blocks and block announcements
// to the remote peer. The goal is to have an async writer that does not lock up
// node internals and at the same time rate limits queued data.
func (p *Peer) broadcastBlocks() {
  2 changes: 1 addition & 1 deletion2  
eth/protocols/eth/dispatcher.go
@@ -224,7 +224,7 @@ func (p *Peer) dispatcher() {
			switch {
			case res.Req == nil:
				// Response arrived with an untracked ID. Since even cancelled
				// requests are tracked until fulfilment, a dangling repsponse
				// requests are tracked until fulfilment, a dangling response
				// means the remote peer implements the protocol badly.
				resOp.fail <- errDanglingResponse

  2 changes: 1 addition & 1 deletion2  
eth/protocols/eth/handler_test.go
@@ -94,7 +94,7 @@ func (b *testBackend) Chain() *core.BlockChain { return b.chain }
func (b *testBackend) TxPool() TxPool          { return b.txpool }

func (b *testBackend) RunPeer(peer *Peer, handler Handler) error {
	// Normally the backend would do peer mainentance and handshakes. All that
	// Normally the backend would do peer maintenance and handshakes. All that
	// is omitted and we will just give control back to the handler.
	return handler(peer)
}
  2 changes: 1 addition & 1 deletion2  
eth/protocols/eth/peer.go
@@ -133,7 +133,7 @@ func (p *Peer) ID() string {
	return p.id
}

// Version retrieves the peer's negoatiated `eth` protocol version.
// Version retrieves the peer's negotiated `eth` protocol version.
func (p *Peer) Version() uint {
	return p.version
}
  2 changes: 1 addition & 1 deletion2  
eth/protocols/snap/handler.go
@@ -504,7 +504,7 @@ func ServiceGetTrieNodesQuery(chain *core.BlockChain, req *GetTrieNodesPacket, s
	var (
		nodes [][]byte
		bytes uint64
		loads int // Trie hash expansions to cound database reads
		loads int // Trie hash expansions to count database reads
	)
	for _, pathset := range req.Paths {
		switch len(pathset) {
  8 changes: 4 additions & 4 deletions8  
eth/protocols/snap/peer.go
@@ -61,12 +61,12 @@ func (p *Peer) ID() string {
	return p.id
}

// Version retrieves the peer's negoatiated `snap` protocol version.
// Version retrieves the peer's negotiated `snap` protocol version.
func (p *Peer) Version() uint {
	return p.version
}

// Log overrides the P2P logget with the higher level one containing only the id.
// Log overrides the P2P logger with the higher level one containing only the id.
func (p *Peer) Log() log.Logger {
	return p.logger
}
@@ -87,7 +87,7 @@ func (p *Peer) RequestAccountRange(id uint64, root common.Hash, origin, limit co
}

// RequestStorageRange fetches a batch of storage slots belonging to one or more
// accounts. If slots from only one accout is requested, an origin marker may also
// accounts. If slots from only one account is requested, an origin marker may also
// be used to retrieve from there.
func (p *Peer) RequestStorageRanges(id uint64, root common.Hash, accounts []common.Hash, origin, limit []byte, bytes uint64) error {
	if len(accounts) == 1 && origin != nil {
@@ -119,7 +119,7 @@ func (p *Peer) RequestByteCodes(id uint64, hashes []common.Hash, bytes uint64) e
}

// RequestTrieNodes fetches a batch of account or storage trie nodes rooted in
// a specificstate trie.
// a specific state trie.
func (p *Peer) RequestTrieNodes(id uint64, root common.Hash, paths []TrieNodePathSet, bytes uint64) error {
	p.logger.Trace("Fetching set of trie nodes", "reqid", id, "root", root, "pathsets", len(paths), "bytes", common.StorageSize(bytes))

  26 changes: 13 additions & 13 deletions26  
eth/protocols/snap/sync.go
@@ -365,15 +365,15 @@ type SyncPeer interface {
	RequestAccountRange(id uint64, root, origin, limit common.Hash, bytes uint64) error

	// RequestStorageRanges fetches a batch of storage slots belonging to one or
	// more accounts. If slots from only one accout is requested, an origin marker
	// more accounts. If slots from only one account is requested, an origin marker
	// may also be used to retrieve from there.
	RequestStorageRanges(id uint64, root common.Hash, accounts []common.Hash, origin, limit []byte, bytes uint64) error

	// RequestByteCodes fetches a batch of bytecodes by hash.
	RequestByteCodes(id uint64, hashes []common.Hash, bytes uint64) error

	// RequestTrieNodes fetches a batch of account or storage trie nodes rooted in
	// a specificstate trie.
	// a specific state trie.
	RequestTrieNodes(id uint64, root common.Hash, paths []TrieNodePathSet, bytes uint64) error

	// Log retrieves the peer's own contextual logger.
@@ -1183,10 +1183,10 @@ func (s *Syncer) assignStorageTasks(success chan *storageResponse, fail chan *st
		}
		if subtask == nil {
			// No large contract required retrieval, but small ones available
			for acccount, root := range task.stateTasks {
				delete(task.stateTasks, acccount)
			for account, root := range task.stateTasks {
				delete(task.stateTasks, account)

				accounts = append(accounts, acccount)
				accounts = append(accounts, account)
				roots = append(roots, root)

				if len(accounts) >= storageSets {
@@ -1486,7 +1486,7 @@ func (s *Syncer) assignBytecodeHealTasks(success chan *bytecodeHealResponse, fai
	}
}

// revertRequests locates all the currently pending reuqests from a particular
// revertRequests locates all the currently pending requests from a particular
// peer and reverts them, rescheduling for others to fulfill.
func (s *Syncer) revertRequests(peer string) {
	// Gather the requests first, revertals need the lock too
@@ -1575,7 +1575,7 @@ func (s *Syncer) revertAccountRequest(req *accountRequest) {
	s.lock.Unlock()

	// If there's a timeout timer still running, abort it and mark the account
	// task as not-pending, ready for resheduling
	// task as not-pending, ready for rescheduling
	req.timeout.Stop()
	if req.task.req == req {
		req.task.req = nil
@@ -1616,7 +1616,7 @@ func (s *Syncer) revertBytecodeRequest(req *bytecodeRequest) {
	s.lock.Unlock()

	// If there's a timeout timer still running, abort it and mark the code
	// retrievals as not-pending, ready for resheduling
	// retrievals as not-pending, ready for rescheduling
	req.timeout.Stop()
	for _, hash := range req.hashes {
		req.task.codeTasks[hash] = struct{}{}
@@ -1657,7 +1657,7 @@ func (s *Syncer) revertStorageRequest(req *storageRequest) {
	s.lock.Unlock()

	// If there's a timeout timer still running, abort it and mark the storage
	// task as not-pending, ready for resheduling
	// task as not-pending, ready for rescheduling
	req.timeout.Stop()
	if req.subTask != nil {
		req.subTask.req = nil
@@ -1743,7 +1743,7 @@ func (s *Syncer) revertBytecodeHealRequest(req *bytecodeHealRequest) {
	s.lock.Unlock()

	// If there's a timeout timer still running, abort it and mark the code
	// retrievals as not-pending, ready for resheduling
	// retrievals as not-pending, ready for rescheduling
	req.timeout.Stop()
	for _, hash := range req.hashes {
		req.task.codeTasks[hash] = struct{}{}
@@ -2035,7 +2035,7 @@ func (s *Syncer) processStorageResponse(res *storageResponse) {
			}
			tr.Commit()
		}
		// Persist the received storage segements. These flat state maybe
		// Persist the received storage segments. These flat state maybe
		// outdated during the sync, but it can be fixed later during the
		// snapshot generation.
		for j := 0; j < len(res.hashes[i]); j++ {
@@ -2170,7 +2170,7 @@ func (s *Syncer) forwardAccountTask(task *accountTask) {
	}
	task.res = nil

	// Persist the received account segements. These flat state maybe
	// Persist the received account segments. These flat state maybe
	// outdated during the sync, but it can be fixed later during the
	// snapshot generation.
	oldAccountBytes := s.accountBytes
@@ -2773,7 +2773,7 @@ func (s *Syncer) onHealByteCodes(peer SyncPeer, id uint64, bytecodes [][]byte) e
}

// onHealState is a callback method to invoke when a flat state(account
// or storage slot) is downloded during the healing stage. The flat states
// or storage slot) is downloaded during the healing stage. The flat states
// can be persisted blindly and can be fixed later in the generation stage.
// Note it's not concurrent safe, please handle the concurrent issue outside.
func (s *Syncer) onHealState(paths [][]byte, value []byte) error {
  2 changes: 1 addition & 1 deletion2  
eth/state_accessor.go
@@ -44,7 +44,7 @@ import (
//        perform Commit or other 'save-to-disk' changes, this should be set to false to avoid
//        storing trash persistently
// - preferDisk: this arg can be used by the caller to signal that even though the 'base' is provided,
//        it would be preferrable to start from a fresh state, if we have it on disk.
//        it would be preferable to start from a fresh state, if we have it on disk.
func (eth *Ethereum) StateAtBlock(block *types.Block, reexec uint64, base *state.StateDB, checkLive bool, preferDisk bool) (statedb *state.StateDB, err error) {
	var (
		current  *types.Block
  10 changes: 5 additions & 5 deletions10  
eth/tracers/api.go
@@ -116,7 +116,7 @@ func (context *chainContext) GetHeader(hash common.Hash, number uint64) *types.H
	return header
}

// chainContext construts the context reader which is used by the evm for reading
// chainContext constructs the context reader which is used by the evm for reading
// the necessary chain context.
func (api *API) chainContext(ctx context.Context) core.ChainContext {
	return &chainContext{api: api, ctx: ctx}
@@ -202,10 +202,10 @@ type blockTraceTask struct {
	statedb *state.StateDB   // Intermediate state prepped for tracing
	block   *types.Block     // Block to trace the transactions from
	rootref common.Hash      // Trie root reference held for this task
	results []*txTraceResult // Trace results procudes by the task
	results []*txTraceResult // Trace results produced by the task
}

// blockTraceResult represets the results of tracing a single block when an entire
// blockTraceResult represents the results of tracing a single block when an entire
// chain is being traced.
type blockTraceResult struct {
	Block  hexutil.Uint64   `json:"block"`  // Block number corresponding to this trace
@@ -563,7 +563,7 @@ func (api *API) StandardTraceBadBlockToFile(ctx context.Context, hash common.Has

// traceBlock configures a new tracer according to the provided configuration, and
// executes all the transactions contained within. The return value will be one item
// per transaction, dependent on the requestd tracer.
// per transaction, dependent on the requested tracer.
func (api *API) traceBlock(ctx context.Context, block *types.Block, config *TraceConfig) ([]*txTraceResult, error) {
	if block.NumberU64() == 0 {
		return nil, errors.New("genesis is not traceable")
@@ -707,7 +707,7 @@ func (api *API) standardTraceBlockToFile(ctx context.Context, block *types.Block
		}
	}
	for i, tx := range block.Transactions() {
		// Prepare the trasaction for un-traced execution
		// Prepare the transaction for un-traced execution
		var (
			msg, _    = tx.AsMessage(signer, block.BaseFee())
			txContext = core.NewEVMTxContext(msg)
  2 changes: 1 addition & 1 deletion2  
eth/tracers/internal/tracetest/calltrace_test.go
@@ -39,7 +39,7 @@ import (
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/tests"

	// Force-load native and js pacakges, to trigger registration
	// Force-load native and js packages, to trigger registration
	_ "github.com/ethereum/go-ethereum/eth/tracers/js"
	_ "github.com/ethereum/go-ethereum/eth/tracers/native"
)
  2 changes: 1 addition & 1 deletion2  
ethclient/ethclient_test.go
@@ -581,7 +581,7 @@ func testCallContract(t *testing.T, client *rpc.Client) {
	if _, err := ec.CallContract(context.Background(), msg, big.NewInt(1)); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	// PendingCallCOntract
	// PendingCallContract
	if _, err := ec.PendingCallContract(context.Background(), msg); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
  2 changes: 1 addition & 1 deletion2  
ethdb/memorydb/memorydb.go
@@ -66,7 +66,7 @@ func NewWithCap(size int) *Database {
}

// Close deallocates the internal map and ensures any consecutive data access op
// failes with an error.
// fails with an error.
func (db *Database) Close() error {
	db.lock.Lock()
	defer db.lock.Unlock()
  2 changes: 1 addition & 1 deletion2  
interfaces.go
@@ -204,7 +204,7 @@ type GasPricer interface {
// FeeHistory provides recent fee market data that consumers can use to determine
// a reasonable maxPriorityFeePerGas value.
type FeeHistory struct {
	OldestBlock  *big.Int     // block coresponding to first response value
	OldestBlock  *big.Int     // block corresponding to first response value
	Reward       [][]*big.Int // list every txs priority fee per block
	BaseFee      []*big.Int   // list of each block's base fee
	GasUsedRatio []float64    // ratio of gas used out of the total available limit
  2 changes: 1 addition & 1 deletion2  
internal/ethapi/api.go
@@ -988,7 +988,7 @@ func newRevertError(result *core.ExecutionResult) *revertError {
	}
}

// revertError is an API error that encompassas an EVM revertal with JSON error
// revertError is an API error that encompasses an EVM revertal with JSON error
// code and a binary data blob.
type revertError struct {
	error
  2 changes: 1 addition & 1 deletion2  
les/downloader/api.go
@@ -125,7 +125,7 @@ type SyncingResult struct {
	Status  ethereum.SyncProgress `json:"status"`
}

// uninstallSyncSubscriptionRequest uninstalles a syncing subscription in the API event loop.
// uninstallSyncSubscriptionRequest uninstalls a syncing subscription in the API event loop.
type uninstallSyncSubscriptionRequest struct {
	c           chan interface{}
	uninstalled chan interface{}
  4 changes: 2 additions & 2 deletions4  
les/downloader/downloader.go
@@ -1625,7 +1625,7 @@ func (d *Downloader) processHeaders(origin uint64, td *big.Int) error {
						log.Warn("Invalid header encountered", "number", chunk[n].Number, "hash", chunk[n].Hash(), "parent", chunk[n].ParentHash, "err", err)
						return fmt.Errorf("%w: %v", errInvalidChain, err)
					}
					// All verifications passed, track all headers within the alloted limits
					// All verifications passed, track all headers within the allotted limits
					if mode == FastSync {
						head := chunk[len(chunk)-1].Number.Uint64()
						if head-rollback > uint64(fsHeaderSafetyNet) {
@@ -1663,7 +1663,7 @@ func (d *Downloader) processHeaders(origin uint64, td *big.Int) error {
			}
			d.syncStatsLock.Unlock()

			// Signal the content downloaders of the availablility of new tasks
			// Signal the content downloaders of the availability of new tasks
			for _, ch := range []chan bool{d.bodyWakeCh, d.receiptWakeCh} {
				select {
				case ch <- true:
  4 changes: 2 additions & 2 deletions4  
les/downloader/downloader_test.go
@@ -653,8 +653,8 @@ func testForkedSync(t *testing.T, protocol uint, mode SyncMode) {
	assertOwnForkedChain(t, tester, testChainBase.len(), []int{chainA.len(), chainB.len()})
}

// Tests that synchronising against a much shorter but much heavyer fork works
// corrently and is not dropped.
// Tests that synchronising against a much shorter but much heavier fork works
// correctly and is not dropped.
func TestHeavyForkedSync66Full(t *testing.T)  { testHeavyForkedSync(t, eth.ETH66, FullSync) }
func TestHeavyForkedSync66Fast(t *testing.T)  { testHeavyForkedSync(t, eth.ETH66, FastSync) }
func TestHeavyForkedSync66Light(t *testing.T) { testHeavyForkedSync(t, eth.ETH66, LightSync) }
  2 changes: 1 addition & 1 deletion2  
les/downloader/queue.go
@@ -872,7 +872,7 @@ func (q *queue) deliver(id string, taskPool map[common.Hash]*types.Header,
		if res, stale, err := q.resultCache.GetDeliverySlot(header.Number.Uint64()); err == nil {
			reconstruct(accepted, res)
		} else {
			// else: betweeen here and above, some other peer filled this result,
			// else: between here and above, some other peer filled this result,
			// or it was indeed a no-op. This should not happen, but if it does it's
			// not something to panic about
			log.Error("Delivery stale", "stale", stale, "number", header.Number.Uint64(), "err", err)
  2 changes: 1 addition & 1 deletion2  
les/fetcher.go
@@ -252,7 +252,7 @@ func (f *lightFetcher) forEachPeer(check func(id enode.ID, p *fetcherPeer) bool)
//   request will be made for header retrieval.
//
// - re-sync trigger
//   If the local chain lags too much, then the fetcher will enter "synnchronise"
//   If the local chain lags too much, then the fetcher will enter "synchronise"
//   mode to retrieve missing headers in batch.
func (f *lightFetcher) mainloop() {
	defer f.wg.Done()
  2 changes: 1 addition & 1 deletion2  
les/flowcontrol/manager.go
@@ -55,7 +55,7 @@ var (
// ClientManager controls the capacity assigned to the clients of a server.
// Since ServerParams guarantee a safe lower estimate for processable requests
// even in case of all clients being active, ClientManager calculates a
// corrigated buffer value and usually allows a higher remaining buffer value
// corrugated buffer value and usually allows a higher remaining buffer value
// to be returned with each reply.
type ClientManager struct {
	clock mclock.Clock
  2 changes: 1 addition & 1 deletion2  
les/odr.go
@@ -126,7 +126,7 @@ const (
// RetrieveTxStatus retrieves the transaction status from the LES network.
// There is no guarantee in the LES protocol that the mined transaction will
// be retrieved back for sure because of different reasons(the transaction
// is unindexed, the malicous server doesn't reply it deliberately, etc).
// is unindexed, the malicious server doesn't reply it deliberately, etc).
// Therefore, unretrieved transactions(UNKNOWN) will receive a certain number
// of retries, thus giving a weak guarantee.
func (odr *LesOdr) RetrieveTxStatus(ctx context.Context, req *light.TxStatusRequest) error {
  2 changes: 1 addition & 1 deletion2  
les/vflux/client/fillset_test.go
@@ -104,7 +104,7 @@ func TestFillSet(t *testing.T) {
	fs.SetTarget(10)
	expWaiting(4, true)
	expNotWaiting()
	// remove all previosly set flags
	// remove all previously set flags
	ns.ForEach(sfTest1, nodestate.Flags{}, func(node *enode.Node, state nodestate.Flags) {
		ns.SetState(node, nodestate.Flags{}, sfTest1, 0)
	})
  2 changes: 1 addition & 1 deletion2  
les/vflux/client/serverpool_test.go
@@ -66,7 +66,7 @@ type ServerPoolTest struct {
	// (accessed from both the main thread and the preNeg callback)
	preNegLock sync.Mutex
	queryWg    *sync.WaitGroup // a new wait group is created each time the simulation is started
	stopping   bool            // stopping avoid callind queryWg.Add after queryWg.Wait
	stopping   bool            // stopping avoid calling queryWg.Add after queryWg.Wait

	cycle, conn, servedConn  int
	serviceCycles, dialCount int
  2 changes: 1 addition & 1 deletion2  
les/vflux/server/balance.go
@@ -623,7 +623,7 @@ func (n *nodeBalance) priorityToBalance(priority int64, capacity uint64) (uint64
	return 0, uint64(-priority)
}

// reducedBalance estimates the reduced balance at a given time in the fututre based
// reducedBalance estimates the reduced balance at a given time in the future based
// on the given balance, the time factor and an estimated average request cost per time ratio
func (n *nodeBalance) reducedBalance(b balance, start mclock.AbsTime, dt time.Duration, capacity uint64, avgReqCost float64) balance {
	// since the costs are applied continuously during the dt time period we calculate
  4 changes: 2 additions & 2 deletions4  
les/vflux/server/balance_test.go
@@ -54,7 +54,7 @@ func newBalanceTestSetup(db ethdb.KeyValueStore, posExp, negExp utils.ValueExpir
	// Initialize and customize the setup for the balance testing
	clock := &mclock.Simulated{}
	setup := newServerSetup()
	setup.clientField = setup.setup.NewField("balancTestClient", reflect.TypeOf(balanceTestClient{}))
	setup.clientField = setup.setup.NewField("balanceTestClient", reflect.TypeOf(balanceTestClient{}))

	ns := nodestate.NewNodeStateMachine(nil, nil, clock, setup.setup)
	if posExp == nil {
@@ -298,7 +298,7 @@ func TestEstimatedPriority(t *testing.T) {
	}
}

func TestPostiveBalanceCounting(t *testing.T) {
func TestPositiveBalanceCounting(t *testing.T) {
	b := newBalanceTestSetup(nil, nil, nil)
	defer b.stop()

  2 changes: 1 addition & 1 deletion2  
les/vflux/server/status.go
@@ -41,7 +41,7 @@ type serverSetup struct {
	activeFlag    nodestate.Flags // Flag is set if the node is active
	inactiveFlag  nodestate.Flags // Flag is set if the node is inactive
	capacityField nodestate.Field // Field contains the capacity of the node
	queueField    nodestate.Field // Field contains the infomration in the priority queue
	queueField    nodestate.Field // Field contains the information in the priority queue
}

// newServerSetup initializes the setup for state machine and returns the flags/fields group.
  2 changes: 1 addition & 1 deletion2  
light/lightchain.go
@@ -397,7 +397,7 @@ func (lc *LightChain) SetCanonical(header *types.Header) error {
//
// The verify parameter can be used to fine tune whether nonce verification
// should be done or not. The reason behind the optional check is because some
// of the header retrieval mechanisms already need to verfy nonces, as well as
// of the header retrieval mechanisms already need to verify nonces, as well as
// because nonces can be verified sparsely, not needing to check each.
//
// In the case of a light chain, InsertHeaderChain also creates and posts light
  4 changes: 2 additions & 2 deletions4  
light/odr_util.go
@@ -272,9 +272,9 @@ func GetBloomBits(ctx context.Context, odr OdrBackend, bit uint, sections []uint
// GetTransaction retrieves a canonical transaction by hash and also returns
// its position in the chain. There is no guarantee in the LES protocol that
// the mined transaction will be retrieved back for sure because of different
// reasons(the transaction is unindexed, the malicous server doesn't reply it
// reasons(the transaction is unindexed, the malicious server doesn't reply it
// deliberately, etc). Therefore, unretrieved transactions will receive a certain
// number of retrys, thus giving a weak guarantee.
// number of retries, thus giving a weak guarantee.
func GetTransaction(ctx context.Context, odr OdrBackend, txHash common.Hash) (*types.Transaction, common.Hash, uint64, uint64, error) {
	r := &TxStatusRequest{Hashes: []common.Hash{txHash}}
	if err := odr.RetrieveTxStatus(ctx, r); err != nil || r.Status[0].Status != core.TxStatusIncluded {
  4 changes: 2 additions & 2 deletions4  
light/postprocess.go
@@ -313,15 +313,15 @@ var (
	BloomTrieTablePrefix = "blt-"
)

// GetBloomTrieRoot reads the BloomTrie root assoctiated to the given section from the database
// GetBloomTrieRoot reads the BloomTrie root associated to the given section from the database
func GetBloomTrieRoot(db ethdb.Database, sectionIdx uint64, sectionHead common.Hash) common.Hash {
	var encNumber [8]byte
	binary.BigEndian.PutUint64(encNumber[:], sectionIdx)
	data, _ := db.Get(append(append(bloomTriePrefix, encNumber[:]...), sectionHead.Bytes()...))
	return common.BytesToHash(data)
}

// StoreBloomTrieRoot writes the BloomTrie root assoctiated to the given section into the database
// StoreBloomTrieRoot writes the BloomTrie root associated to the given section into the database
func StoreBloomTrieRoot(db ethdb.Database, sectionIdx uint64, sectionHead, root common.Hash) {
	var encNumber [8]byte
	binary.BigEndian.PutUint64(encNumber[:], sectionIdx)
  2 changes: 1 addition & 1 deletion2  
light/trie.go
@@ -153,7 +153,7 @@ func (t *odrTrie) TryDelete(key []byte) error {
	})
}

// TryDeleteACcount abstracts an account deletion from the trie.
// TryDeleteAccount abstracts an account deletion from the trie.
func (t *odrTrie) TryDeleteAccount(key []byte) error {
	key = crypto.Keccak256(key)
	return t.do(key, func() error {
  2 changes: 1 addition & 1 deletion2  
light/txpool.go
@@ -71,7 +71,7 @@ type TxPool struct {
	eip2718  bool // Fork indicator whether we are in the eip2718 stage.
}

// TxRelayBackend provides an interface to the mechanism that forwards transacions
// TxRelayBackend provides an interface to the mechanism that forwards transactions
// to the ETH network. The implementations of the functions should be non-blocking.
//
// Send instructs backend to forward new transactions
  2 changes: 1 addition & 1 deletion2  
metrics/gauge_float64_test.go
@@ -2,7 +2,7 @@ package metrics

import "testing"

func BenchmarkGuageFloat64(b *testing.B) {
func BenchmarkGaugeFloat64(b *testing.B) {
	g := NewGaugeFloat64()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
  2 changes: 1 addition & 1 deletion2  
metrics/gauge_test.go
@@ -5,7 +5,7 @@ import (
	"testing"
)

func BenchmarkGuage(b *testing.B) {
func BenchmarkGauge(b *testing.B) {
	g := NewGauge()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
  2 changes: 1 addition & 1 deletion2  
metrics/prometheus/prometheus.go
@@ -36,7 +36,7 @@ func Handler(reg metrics.Registry) http.Handler {
		})
		sort.Strings(names)

		// Aggregate all the metris into a Prometheus collector
		// Aggregate all the metrics into a Prometheus collector
		c := newCollector()

		for _, name := range names {
  2 changes: 1 addition & 1 deletion2  
miner/unconfirmed_test.go
@@ -74,7 +74,7 @@ func TestUnconfirmedShifts(t *testing.T) {
	if n := pool.blocks.Len(); n != int(limit)/2 {
		t.Errorf("unconfirmed count mismatch: have %d, want %d", n, limit/2)
	}
	// Try to shift all the remaining blocks out and verify emptyness
	// Try to shift all the remaining blocks out and verify emptiness
	pool.Shift(start + 2*uint64(limit))
	if n := pool.blocks.Len(); n != 0 {
		t.Errorf("unconfirmed count mismatch: have %d, want %d", n, 0)
  2 changes: 1 addition & 1 deletion2  
miner/worker_test.go
@@ -494,7 +494,7 @@ func testAdjustInterval(t *testing.T, chainConfig *params.ChainConfig, engine co
	}
	w.start()

	time.Sleep(time.Second) // Ensure two tasks have been summitted due to start opt
	time.Sleep(time.Second) // Ensure two tasks have been submitted due to start opt
	atomic.StoreUint32(&start, 1)

	w.setRecommitInterval(3 * time.Second)
  6 changes: 3 additions & 3 deletions6  
mobile/accounts.go
@@ -212,10 +212,10 @@ func (ks *KeyStore) ImportECDSAKey(key []byte, passphrase string) (account *Acco

// ImportPreSaleKey decrypts the given Ethereum presale wallet and stores
// a key file in the key directory. The key file is encrypted with the same passphrase.
func (ks *KeyStore) ImportPreSaleKey(keyJSON []byte, passphrase string) (ccount *Account, _ error) {
	account, err := ks.keystore.ImportPreSaleKey(common.CopyBytes(keyJSON), passphrase)
func (ks *KeyStore) ImportPreSaleKey(keyJSON []byte, passphrase string) (account *Account, _ error) {
	acc, err := ks.keystore.ImportPreSaleKey(common.CopyBytes(keyJSON), passphrase)
	if err != nil {
		return nil, err
	}
	return &Account{account}, nil
	return &Account{acc}, nil
}
  2 changes: 1 addition & 1 deletion2  
mobile/init.go
@@ -14,7 +14,7 @@
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

// Contains initialization code for the mbile library.
// Contains initialization code for the mobile library.

package geth

  2 changes: 1 addition & 1 deletion2  
node/rpcstack_test.go
@@ -100,7 +100,7 @@ func TestWebsocketOrigins(t *testing.T) {
			expFail: []string{
				"test",                                // no scheme, required by spec
				"http://test",                         // wrong scheme
				"http://test.foo", "https://a.test.x", // subdomain variatoins
				"http://test.foo", "https://a.test.x", // subdomain variations
				"http://testx:8540", "https://xtest:8540"},
		},
		// ip tests
  2 changes: 1 addition & 1 deletion2  
p2p/discover/v4_udp.go
@@ -525,7 +525,7 @@ func (t *UDPv4) readLoop(unhandled chan<- ReadPacket) {
			t.log.Debug("Temporary UDP read error", "err", err)
			continue
		} else if err != nil {
			// Shut down the loop for permament errors.
			// Shut down the loop for permanent errors.
			if !errors.Is(err, io.EOF) {
				t.log.Debug("UDP read error", "err", err)
			}
  2 changes: 1 addition & 1 deletion2  
p2p/discover/v5_udp.go
@@ -625,7 +625,7 @@ func (t *UDPv5) readLoop() {
			t.log.Debug("Temporary UDP read error", "err", err)
			continue
		} else if err != nil {
			// Shut down the loop for permament errors.
			// Shut down the loop for permanent errors.
			if !errors.Is(err, io.EOF) {
				t.log.Debug("UDP read error", "err", err)
			}
  6 changes: 3 additions & 3 deletions6  
p2p/msgrate/msgrate.go
@@ -111,7 +111,7 @@ const tuningImpact = 0.25
// local link is saturated. In that case, the live measurements will force us
// to reduce request sizes until the throughput gets stable.
//
// Lastly, message rate measurements allows us to detect if a peer is unsuaully
// Lastly, message rate measurements allows us to detect if a peer is unusually
// slow compared to other peers, in which case we can decide to keep it around
// or free up the slot so someone closer.
//
@@ -127,7 +127,7 @@ type Tracker struct {
	// in their sizes.
	//
	// Callers of course are free to use the item counter as a byte counter if
	// or when their protocol of choise if capped by bytes instead of items.
	// or when their protocol of choice if capped by bytes instead of items.
	// (eg. eth.getHeaders vs snap.getAccountRange).
	capacity map[uint64]float64

@@ -157,7 +157,7 @@ func NewTracker(caps map[uint64]float64, rtt time.Duration) *Tracker {
}

// Capacity calculates the number of items the peer is estimated to be able to
// retrieve within the alloted time slot. The method will round up any division
// retrieve within the allotted time slot. The method will round up any division
// errors and will add an additional overestimation ratio on top. The reason for
// overshooting the capacity is because certain message types might not increase
// the load proportionally to the requested items, so fetching a bit more might
  2 changes: 1 addition & 1 deletion2  
p2p/tracker/tracker.go
@@ -121,7 +121,7 @@ func (t *Tracker) Track(peer string, version uint, reqCode uint64, resCode uint6
}

// clean is called automatically when a preset time passes without a response
// being dleivered for the first network request.
// being delivered for the first network request.
func (t *Tracker) clean() {
	t.lock.Lock()
	defer t.lock.Unlock()
  2 changes: 1 addition & 1 deletion2  
rpc/server.go
@@ -160,7 +160,7 @@ type PeerInfo struct {
	// Address of client. This will usually contain the IP address and port.
	RemoteAddr string

	// Addditional information for HTTP and WebSocket connections.
	// Additional information for HTTP and WebSocket connections.
	HTTP struct {
		// Protocol version, i.e. "HTTP/1.1". This is not set for WebSocket.
		Version string
  8 changes: 4 additions & 4 deletions8  
signer/rules/rules_test.go
@@ -44,7 +44,7 @@ Three things can happen:
3. Anything else; other return values [*], method not implemented or exception occurred during processing. This means
that the operation will continue to manual processing, via the regular UI method chosen by the user.
[*] Note: Future version of the ruleset may use more complex json-based returnvalues, making it possible to not
[*] Note: Future version of the ruleset may use more complex json-based return values, making it possible to not
only respond Approve/Reject/Manual, but also modify responses. For example, choose to list only one, but not all
accounts in a list-request. The points above will continue to hold for non-json based responses ("Approve"/"Reject").
@@ -242,7 +242,7 @@ func (d *dummyUI) OnApprovedTx(tx ethapi.SignTransactionResult) {
func (d *dummyUI) OnSignerStartup(info core.StartupInfo) {
}

//TestForwarding tests that the rule-engine correctly dispatches requests to the next caller
// TestForwarding tests that the rule-engine correctly dispatches requests to the next caller
func TestForwarding(t *testing.T) {
	js := ""
	ui := &dummyUI{make([]string, 0)}
@@ -434,7 +434,7 @@ func dummyTx(value hexutil.Big) *core.SignTxRequest {
			Gas:      gas,
		},
		Callinfo: []apitypes.ValidationInfo{
			{Typ: "Warning", Message: "All your base are bellong to us"},
			{Typ: "Warning", Message: "All your base are belong to us"},
		},
		Meta: core.Metadata{Remote: "remoteip", Local: "localip", Scheme: "inproc"},
	}
@@ -536,7 +536,7 @@ func (d *dontCallMe) OnApprovedTx(tx ethapi.SignTransactionResult) {
	d.t.Fatalf("Did not expect next-handler to be called")
}

//TestContextIsCleared tests that the rule-engine does not retain variables over several requests.
// TestContextIsCleared tests that the rule-engine does not retain variables over several requests.
// if it does, that would be bad since developers may rely on that to store data,
// instead of using the disk-based data storage
func TestContextIsCleared(t *testing.T) {
  2 changes: 1 addition & 1 deletion2  
signer/storage/aes_gcm_storage.go
@@ -143,7 +143,7 @@ func (s *AESEncryptedStorage) writeEncryptedStorage(creds map[string]storedCrede

// encrypt encrypts plaintext with the given key, with additional data
// The 'additionalData' is used to place the (plaintext) KV-store key into the V,
// to prevent the possibility to alter a K, or swap two entries in the KV store with eachother.
// to prevent the possibility to alter a K, or swap two entries in the KV store with each other.
func encrypt(key []byte, plaintext []byte, additionalData []byte) ([]byte, []byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
  2 changes: 1 addition & 1 deletion2  
trie/hasher.go
@@ -191,7 +191,7 @@ func (h *hasher) hashData(data []byte) hashNode {
}

// proofHash is used to construct trie proofs, and returns the 'collapsed'
// node (for later RLP encoding) aswell as the hashed node -- unless the
// node (for later RLP encoding) as well as the hashed node -- unless the
// node is smaller than 32 bytes, in which case it will be returned as is.
// This method does not do anything on value- or hash-nodes.
func (h *hasher) proofHash(original node) (collapsed, hashed node) {
  32 changes: 16 additions & 16 deletions32  
trie/proof_test.go
@@ -205,7 +205,7 @@ func TestRangeProofWithNonExistentProof(t *testing.T) {
		proof := memorydb.New()

		// Short circuit if the decreased key is same with the previous key
		first := decreseKey(common.CopyBytes(entries[start].k))
		first := decreaseKey(common.CopyBytes(entries[start].k))
		if start != 0 && bytes.Equal(first, entries[start-1].k) {
			continue
		}
@@ -214,7 +214,7 @@ func TestRangeProofWithNonExistentProof(t *testing.T) {
			continue
		}
		// Short circuit if the increased key is same with the next key
		last := increseKey(common.CopyBytes(entries[end-1].k))
		last := increaseKey(common.CopyBytes(entries[end-1].k))
		if end != len(entries) && bytes.Equal(last, entries[end].k) {
			continue
		}
@@ -274,7 +274,7 @@ func TestRangeProofWithInvalidNonExistentProof(t *testing.T) {

	// Case 1
	start, end := 100, 200
	first := decreseKey(common.CopyBytes(entries[start].k))
	first := decreaseKey(common.CopyBytes(entries[start].k))

	proof := memorydb.New()
	if err := trie.Prove(first, 0, proof); err != nil {
@@ -297,7 +297,7 @@ func TestRangeProofWithInvalidNonExistentProof(t *testing.T) {

	// Case 2
	start, end = 100, 200
	last := increseKey(common.CopyBytes(entries[end-1].k))
	last := increaseKey(common.CopyBytes(entries[end-1].k))
	proof = memorydb.New()
	if err := trie.Prove(entries[start].k, 0, proof); err != nil {
		t.Fatalf("Failed to prove the first node %v", err)
@@ -343,7 +343,7 @@ func TestOneElementRangeProof(t *testing.T) {

	// One element with left non-existent edge proof
	start = 1000
	first := decreseKey(common.CopyBytes(entries[start].k))
	first := decreaseKey(common.CopyBytes(entries[start].k))
	proof = memorydb.New()
	if err := trie.Prove(first, 0, proof); err != nil {
		t.Fatalf("Failed to prove the first node %v", err)
@@ -358,7 +358,7 @@ func TestOneElementRangeProof(t *testing.T) {

	// One element with right non-existent edge proof
	start = 1000
	last := increseKey(common.CopyBytes(entries[start].k))
	last := increaseKey(common.CopyBytes(entries[start].k))
	proof = memorydb.New()
	if err := trie.Prove(entries[start].k, 0, proof); err != nil {
		t.Fatalf("Failed to prove the first node %v", err)
@@ -373,7 +373,7 @@ func TestOneElementRangeProof(t *testing.T) {

	// One element with two non-existent edge proofs
	start = 1000
	first, last = decreseKey(common.CopyBytes(entries[start].k)), increseKey(common.CopyBytes(entries[start].k))
	first, last = decreaseKey(common.CopyBytes(entries[start].k)), increaseKey(common.CopyBytes(entries[start].k))
	proof = memorydb.New()
	if err := trie.Prove(first, 0, proof); err != nil {
		t.Fatalf("Failed to prove the first node %v", err)
@@ -641,9 +641,9 @@ func TestSameSideProofs(t *testing.T) {
	sort.Sort(entries)

	pos := 1000
	first := decreseKey(common.CopyBytes(entries[pos].k))
	first = decreseKey(first)
	last := decreseKey(common.CopyBytes(entries[pos].k))
	first := decreaseKey(common.CopyBytes(entries[pos].k))
	first = decreaseKey(first)
	last := decreaseKey(common.CopyBytes(entries[pos].k))

	proof := memorydb.New()
	if err := trie.Prove(first, 0, proof); err != nil {
@@ -657,9 +657,9 @@ func TestSameSideProofs(t *testing.T) {
		t.Fatalf("Expected error, got nil")
	}

	first = increseKey(common.CopyBytes(entries[pos].k))
	last = increseKey(common.CopyBytes(entries[pos].k))
	last = increseKey(last)
	first = increaseKey(common.CopyBytes(entries[pos].k))
	last = increaseKey(common.CopyBytes(entries[pos].k))
	last = increaseKey(last)

	proof = memorydb.New()
	if err := trie.Prove(first, 0, proof); err != nil {
@@ -765,7 +765,7 @@ func TestEmptyRangeProof(t *testing.T) {
	}
	for _, c := range cases {
		proof := memorydb.New()
		first := increseKey(common.CopyBytes(entries[c.pos].k))
		first := increaseKey(common.CopyBytes(entries[c.pos].k))
		if err := trie.Prove(first, 0, proof); err != nil {
			t.Fatalf("Failed to prove the first node %v", err)
		}
@@ -904,7 +904,7 @@ func mutateByte(b []byte) {
	}
}

func increseKey(key []byte) []byte {
func increaseKey(key []byte) []byte {
	for i := len(key) - 1; i >= 0; i-- {
		key[i]++
		if key[i] != 0x0 {
@@ -914,7 +914,7 @@ func increseKey(key []byte) []byte {
	return key
}

func decreseKey(key []byte) []byte {
func decreaseKey(key []byte) []byte {
	for i := len(key) - 1; i >= 0; i-- {
		key[i]--
		if key[i] != 0xff {
  2 changes: 1 addition & 1 deletion2  
trie/secure_trie_test.go
@@ -121,7 +121,7 @@ func TestStateTrieConcurrency(t *testing.T) {
	for i := 0; i < threads; i++ {
		tries[i] = trie.Copy()
	}
	// Start a batch of goroutines interactng with the trie
	// Start a batch of goroutines interacting with the trie
	pend := new(sync.WaitGroup)
	pend.Add(threads)
	for i := 0; i < threads; i++ {
