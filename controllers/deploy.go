package controllers

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os/exec"
)

var  (
	deployEntryType = "deploy"
)

func (d *DeployController) AliCode()  {
	d.JsonOK(deployEntryType, StringMap{"result": "deploy ok"}, true)
	ReLaunch("/app/images.sh")
}

func (d *DeployController) Github()  {
	matched, _ := VerifySignature(d)
	if !matched {
		err := "Signatures did not match"
		d.JsonError(deployEntryType, err, StringMap{"result": err}, true)
		log.Warn(err)
		return
	}
	log.Info("Signatures is matched ~")
	//return 200 first
	d.JsonOK(deployEntryType, StringMap{"result": "deploy ok"}, true)
	ReLaunch("/app/w.sh")

}

func VerifySignature(d *DeployController) (bool, error) {
	PayloadBody, err := ioutil.ReadAll(d.Ctx.Request.Body)
	if err != nil {
		return false, err
	}
	// Get Header with X-Hub-Signature
	XHubSignature := d.Ctx.Request.Header.Get("X-Hub-Signature")
	signature := getSha1Code(PayloadBody)
	log.Info(signature)
	return XHubSignature == signature, nil
}


func getSha1Code(payloadBody []byte) string {
	h := hmac.New(sha1.New, []byte("hongfeng"))
	h.Write(payloadBody)
	return "sha1=" + hex.EncodeToString(h.Sum(nil))
}



func ReLaunch(cmdStr string) {
	cmd := exec.Command("sh", cmdStr)
	err := cmd.Start()
	if err != nil {
		log.Fatal(err.Error())
	}
	err = cmd.Wait()
}