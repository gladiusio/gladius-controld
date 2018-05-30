package crypto

// Sign a hash
func Sign(passphrase string) (*keystore.Key, Error) {
	var pathTemp string = viper.GetString("DirWallet")
	ks := keystore.NewKeyStore(
		pathTemp,
		keystore.LightScryptN,
		keystore.LightScryptP)

	key, err := keystore.DecryptKey(ks.Export(ks.Wallets()[index].Accounts()[0], ))

}
