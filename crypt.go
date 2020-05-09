//開発環境用、証明書、鍵関係作成ファイル

package main

import(
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"net"
	"os"
	"time"
)

func main1() {
	max:=new(big.Int).Lsh(big.NewInt(1),128)
	serialNumber,_:=rand.Int(rand.Reader ,max)

	subject:=pkix.Name{
		Organization:[]string{"ManningPublicationsCo."},
		OrganizationalUnit:[]string{"Books"},
		CommonName:"GoWebProgramming",
	}

	//証明書のテンプレート
	template:=x509.Certificate{
		SerialNumber:serialNumber,
		Subject:subject,
		NotBefore:time.Now(),
		NotAfter:time.Now().Add(365*24*time.Hour),
		KeyUsage:x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:[]x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		IPAddresses:[]net.IP{net.ParseIP("127.0.0.1")},
	}

	//rsaライブラリで秘密鍵生成
    pk,_:=rsa.GenerateKey(rand.Reader,2048)
	derBytes,_:=x509.CreateCertificate(rand.Reader ,&template ,&template ,&pk.PublicKey ,pk)

	//cert.pem（SSL証明書）作成
	certOut,_:=os.Create("cert.pem")
	pem.Encode(certOut ,&pem.Block{Type:"CERTIFICATE" ,Bytes:derBytes})
	certOut.Close()

	//key.pemに保存
	keyOut,_:=os.Create("key.pem")
	pem.Encode(keyOut ,&pem.Block{Type:"RSAPRIVATEKEY" ,Bytes:x509.MarshalPKCS1PrivateKey(pk)})
	keyOut.Close()
}


