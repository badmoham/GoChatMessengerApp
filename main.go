package main

import (
	"image/color"
	"log"
	"os"

	"GoChatDesktopClient/server"
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

func main() {
	go func() {
		window := new(app.Window)
		err := run(window)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func run(window *app.Window) error {
	theme := material.NewTheme()
	var ops op.Ops
	var signUpBtn widget.Clickable
	var signInBtn widget.Clickable
	var phoneNumberInput widget.Editor
	var passwordInput widget.Editor
	elements := []layout.FlexChild{
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				label := material.H1(theme, "GoChat")
				label.Color = color.NRGBA{R: 2, G: 136, B: 209, A: 255}
				return label.Layout(gtx)
			})
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				label := material.H2(theme, "welcome!")
				label.Color = color.NRGBA{R: 144, G: 202, B: 249, A: 255}
				label.Alignment = text.Middle
				return label.Layout(gtx)
			})
		}),
		layout.Rigid(layout.Spacer{Height: unit.Dp(80)}.Layout),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return material.Editor(theme, &phoneNumberInput, "enter your phone number here ...").Layout(gtx)
			})
		}),
		layout.Rigid(layout.Spacer{Height: unit.Dp(10)}.Layout),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return material.Editor(theme, &passwordInput, "enter your password here ...").Layout(gtx)
			})
		}),
		layout.Rigid(layout.Spacer{Height: unit.Dp(20)}.Layout),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return material.Button(theme, &signInBtn, "Sign In").Layout(gtx)
			})
		}),
		layout.Rigid(layout.Spacer{Height: unit.Dp(20)}.Layout),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return material.Button(theme, &signUpBtn, "Sign Up").Layout(gtx)
			})
		}),
	}

	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)

			if signUpBtn.Clicked(gtx) {
				if err := server.SignUp(phoneNumberInput.Text(), passwordInput.Text()); err != nil {
					// TODO: show error
				} else {
					// TODO: enter chat list area
				}
			}
			if signInBtn.Clicked(gtx) {
				if err := server.SignIn(phoneNumberInput.Text(), passwordInput.Text()); err != nil {
					// TODO: show error
				} else {
					// TODO: enter chat list area
				}
			}

			// rendering graphical elements
			layout.Flex{
				Axis: layout.Vertical, // Stack children vertically
			}.Layout(gtx, elements...)
			e.Frame(gtx.Ops)
		}
	}
}
