/**
 * Created by GoLand.
 * @User: yanzhipeng
 * @Date: 2020-07-12
 * @Time: 16:43
 * @Project_Name: cloudSmile
 */

package main

import (
	"fmt"
	"github.com/go-vgo/robotgo"
)


func alert() {
	// show Alert Window
	abool := robotgo.ShowAlert("hello", "robotgo")
	if abool == 0 {
		fmt.Println("ok@@@", "ok")
	}
	robotgo.ShowAlert("hello", "robotgo", "Ok", "Cancel")

}

func get() {
	// get the current process id
	pid := robotgo.GetPID()
	fmt.Println("pid----", pid)

	// get current Window Active
	mdata := robotgo.GetActive()

	// get current Window Handle
	hwnd := robotgo.GetHandle()
	fmt.Println("hwnd---", hwnd)

	// get current Window Handle
	bhwnd := robotgo.GetBHandle()
	fmt.Println("bhwnd---", bhwnd)

	// get current Window title
	title := robotgo.GetTitle()
	fmt.Println("title-----", title)

	// set Window Active
	robotgo.SetActive(mdata)
}

func findIds() {
	// find the process id by the process name
	fpid, err := robotgo.FindIds("Google")
	if err == nil {
		fmt.Println("pids...", fpid)
		if len(fpid) > 0 {
			robotgo.ActivePID(fpid[0])

			tl := robotgo.GetTitle(fpid[0])
			fmt.Println("pid[0] title is: ", tl)

			x, y, w, h := robotgo.GetBounds(fpid[0])
			fmt.Println("GetBounds is: ", x, y, w, h)

			// Windows
			// hwnd := robotgo.FindWindow("google")
			// hwnd := robotgo.GetHWND()
			robotgo.MinWindow(fpid[0])
			robotgo.MaxWindow(fpid[0])
			robotgo.CloseWindow(fpid[0])

			robotgo.Kill(fpid[0])
		}
	}
}

func active() {
	robotgo.ActivePID(100)
	// robotgo.Sleep(2)
	robotgo.ActiveName("code")
	robotgo.Sleep(1)
	robotgo.ActiveName("chrome")
}

func findName() {
	// find the process name by the process id
	name, err := robotgo.FindName(100)
	if err == nil {
		fmt.Println("name: ", name)
	}

	// find the all process name
	names, err := robotgo.FindNames()
	if err == nil {
		fmt.Println("name: ", names)
	}

	p, err := robotgo.FindPath(100)
	if err == nil {
		fmt.Println("path: ", p)
	}
}

func ps() {
	// determine whether the process exists
	isExist, err := robotgo.PidExists(100)
	if err == nil && isExist {
		fmt.Println("pid exists is", isExist)

		robotgo.Kill(100)
	}

	// get the all process id
	pids, err := robotgo.Pids()
	if err == nil {
		fmt.Println("pids: ", pids)
	}

	// get the all process struct
	ps, err := robotgo.Process()
	if err == nil {
		fmt.Println("process: ", ps)
	}
}

func window() {
	////////////////////////////////////////////////////////////////////////////////
	// Window Handle
	////////////////////////////////////////////////////////////////////////////////

	alert()
	//
	get()

	findIds()
	active()

	findName()
	//
	ps()

	// close current Window
	robotgo.CloseWindow()
}

func main() {
	window()
}