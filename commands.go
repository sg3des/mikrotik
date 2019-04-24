package mikrotik

import (
	"fmt"
	"log"

	routeros "gopkg.in/routeros.v2"
)

// ===============================
//
// Mikrotik CMDs
//
// ===============================

//Run one line command on mikrotik by api
func (mik *Mikrotik) Run(cmd string) (*routeros.Reply, error) {
	mik.connMutex.Lock()
	defer mik.connMutex.Unlock()

	re, err := mik.Conn.Run(cmd)
	if err != nil {
		mik.Conn.Run("")
	}
	return re, err
}

//RunArgs run many line command on mikrotik by api
func (mik *Mikrotik) RunArgs(cmd string, args ...string) (*routeros.Reply, error) {
	mik.connMutex.Lock()
	defer mik.connMutex.Unlock()

	re, err := mik.Conn.RunArgs(append([]string{cmd}, args...))

	if err != nil {
		mik.Conn.Run("")
	}
	return re, err
}

//RunMarshal - run command and marshal response to interface struct
func (mik *Mikrotik) RunMarshal(cmd string, v interface{}) error {
	re, err := mik.Run(cmd)
	if err != nil {
		return err
	}

	return mik.ParseResponce(re, v)
}

// ParseResponce - Parse the mikrotik's reply into the interface using reflect.
func (mik *Mikrotik) ParseResponce(re *routeros.Reply, v interface{}) error {
	for _, resp := range re.Re {
		if mik.debug {
			log.Println(resp)
		}

		if err := ValuesFrom(resp.Map).To(v); err != nil {
			return err
		}
	}

	return nil
}

//Print returns all items by apipath and marshal it to passed structure
func (mik *Mikrotik) Print(apipath string, v interface{}) error {
	return mik.RunMarshal(apipath, v)
}

//Add item from passed struct to apipath
func (mik *Mikrotik) Add(apipath string, v interface{}) error {
	re, err := mik.RunArgs(apipath, ToArgs(v)...)
	if err != nil {
		return err
	}

	if id, ok := re.Done.Map["ret"]; ok {
		SetID(v, id)
	}

	return nil
}

//Set value of item by id
func (mik *Mikrotik) Set(apipath, id string, v interface{}) error {
	args := append([]string{"=.id=" + id}, ToArgs(v)...)
	_, err := mik.RunArgs(apipath, args...)
	return err
}

//SetOne set value to one field
func (mik *Mikrotik) SetOne(apipath, name, value string) error {
	_, err := mik.RunArgs(apipath, fmt.Sprintf("=%s=%s", name, value))
	return err
}

//Remove item by id
func (mik *Mikrotik) Remove(apipath, id string) error {
	_, err := mik.RunArgs(apipath, "=.id="+id)
	return err
}

//Enable item by id
func (mik *Mikrotik) Enable(apipath, id string) error {
	_, err := mik.RunArgs(apipath, "=.id="+id)
	return err
}

//Disable item by id
func (mik *Mikrotik) Disable(apipath, id string) error {
	_, err := mik.RunArgs(apipath, "=.id="+id)
	return err
}

//Comment - add comment to item by id
func (mik *Mikrotik) Comment(apipath, id, comment string) error {
	_, err := mik.RunArgs(apipath, "=.id="+id, "=comment="+comment)
	return err
}

// Debug activates the debug mode on the library
func (mik *Mikrotik) Debug(debug bool) {
	mik.debug = debug
}

// Close closes the connection with the Mikrotik
func (mik *Mikrotik) Close() {
	mik.Conn.Close()
}

// Delay creates a sleep mode for the mikrotik
func (mik *Mikrotik) Delay(time string) error {
	_, err := mik.RunArgs("/delay", "=delay-time="+time)

	return err
}

// Reboot reboots the mikrotik
func (mik *Mikrotik) Reboot() error {
	_, err := mik.RunArgs("/system/reboot")

	return err
}

// Upgrade upgrades the mikrotik
func (mik *Mikrotik) Upgrade() error {
	_, err := mik.RunArgs("/system/routerboard/upgrade")

	return err
}

// ===============================
//
// Basic CMDs
//
// ===============================

// The cmd struct represents most of the mikrotik's struct including commands like Add, Print, Find, Comment and whatever.
type cmd struct {
	mikrotik *Mikrotik
	path     string
}

// Print simply calls the mikrotik's Print for the commands that use cmd struct.
func (c *cmd) Print(v interface{}) error {
	return c.mikrotik.Print(c.path+"/print", v)
}

// Find runs a print using the where argument searching for a precise attribute.
func (c *cmd) Find(where string, v interface{}) error {
	re, err := c.mikrotik.RunArgs(c.path+"/print", "?"+where)
	if err != nil {
		return err
	}

	return c.mikrotik.ParseResponce(re, v)
}

// PrintWhere prints all the information about a specific entity using its ID
func (c *cmd) PrintWhere(id string, v interface{}) error {
	re, err := c.mikrotik.RunArgs(c.path+"/print", "where", "?.id="+id)
	if err != nil {
		return err
	}

	return c.mikrotik.ParseResponce(re, v)
}

// Add an entity to the mikrotik
func (c *cmd) Add(v interface{}) error {
	return c.mikrotik.Add(c.path+"/add", v)
}

//Set value of item by id
func (c *cmd) Set(id string, v interface{}) error {
	return c.mikrotik.Set(c.path+"/set", id, v)
}

//Remove item by id
func (c *cmd) Remove(id string) error {
	return c.mikrotik.Remove(c.path+"/remove", id)
}

//Enable item by id
func (c *cmd) Enable(id string) error {
	return c.mikrotik.Enable(c.path+"/enable", id)
}

//Disable item by id
func (c *cmd) Disable(id string) error {
	return c.mikrotik.Disable(c.path+"/disable", id)
}

//Comment - add comment to item by id
func (c *cmd) Comment(id, comment string) error {
	return c.mikrotik.Comment(c.path+"/comment", id, comment)
}

// ===============================
//
// Configuration CMDs
//
// ===============================

// The cfg struct represents the basic commands for configurations that use only Print and Set.
type cfg struct {
	mikrotik *Mikrotik
	path     string
}

// Print simply calls the mikrotik's Print for the commands that use cfg struct.
func (c *cfg) Print(v interface{}) error {
	return c.mikrotik.Print(c.path+"/print", v)
}

//Set value of item by id
func (c *cfg) Set(name, value string) error {
	return c.mikrotik.SetOne(c.path+"/set", name, value)
}

// ===============================
//
// SSH CMDs
//
// ===============================

// The struct sshcmds includes all the methods for the SSH keys managed in System/User
type sshcmds struct {
	mikrotik *Mikrotik
	path     string
}

// Print simply calls the mikrotik's Print for the commands that use sshcmd struct.
func (ssh *sshcmds) Print(v interface{}) error {
	return ssh.mikrotik.Print(ssh.path+"/print", v)
}

//Remove item by id
func (ssh *sshcmds) Remove(id string) error {
	return ssh.mikrotik.Remove(ssh.path+"/remove", id)
}

// Find runs a print using the where argument searching for a precise attribute.
func (ssh *sshcmds) Find(where string, v interface{}) error {
	re, err := ssh.mikrotik.RunArgs(ssh.path+"/print", "?"+where)
	if err != nil {
		return err
	}

	return ssh.mikrotik.ParseResponce(re, v)
}

// Import imports the specified SSH key into the specified user
func (ssh *sshcmds) Import(user, keyfile string) error {
	_, err := ssh.mikrotik.RunArgs(ssh.path+"/import", "=public-key-file="+keyfile, "=user="+user)

	return err
}
