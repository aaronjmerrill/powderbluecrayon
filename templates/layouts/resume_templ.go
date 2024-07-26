// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.747
package layouts

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "powderbluecrayon/templates/components"

func Resume() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<section class=\"mt-3 md:mt-5\"><div><nav class=\"flow-root\"><a href=\"/static/resume.pdf\" download><button class=\"float-end btn btn-outline\">Get my resume</button></a></nav><div class=\"text-center mt-7 md:mt-16\"><div class=\"flex justify-center mb-16 rounded-full\"><img src=\"static/images/profile.png\" alt=\"Image\" class=\"rounded-full w-96\"></div><h6 class=\"font-medium text-lg md:text-2xl uppercase mb-8\">Aaron Merrill</h6><h1 class=\"font-normal text-4xl md:text-7xl leading-none mb-8\">Senior Developer / Cloud Engineer</h1><p class=\"font-normal text-md md:text-xl mb-16\">I have a passion for software. I enjoy creating tools that make life easier for others.</p><a href=\"mailto:ajm@powderbluecrayon.com\" class=\"btn btn-primary btn-lg\">Hire me</a></div></div></section><section class=\"mt-7 md:mt-16\"><div class=\"container max-w-screen-xl mx-auto px-4\"><div class=\"grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6\"><div class=\"custom-card\"><div class=\"w-20 py-6 flex justify-center bg-base-300 rounded-md mb-4\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = components.Activity().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div><h4 class=\"font-medium text-lg mb-4\">Fullstack webapp developer</h4><p class=\"font-normal text-md\">I have led multiple projects to successful launches. I've learned new technologies as needed. And kept up with the latest features of existing languages I've been using</p></div><div class=\"custom-card\"><div class=\"w-20 py-6 flex justify-center bg-base-300 rounded-md mb-4\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = components.Box().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div><h4 class=\"font-medium text-lg mb-4\">Azure cloud engineer</h4><p class=\"font-normal text-md\">I've architected cloud solutions and implemented them with best practices recommended from Microsoft. Working on certifications</p></div><div class=\"custom-card\"><div class=\"w-20 py-6 flex justify-center bg-base-300 rounded-md mb-4\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = components.Coffee().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div><h4 class=\"font-medium text-lg mb-4\">Passion projects</h4><p class=\"font-normal text-md\">I'm passionate about htmx and Go and I wrote this website to gain experience in those tools. I also run several FOSS applications in docker containers on my personal server</p></div></div></div></section><section class=\"mt-7 md:mt-16\"><div class=\"container max-w-screen-xl mx-auto px-4\"><h1 class=\"font-medium text-3xl md:text-4xl mb-5\">Experience</h1><div class=\"flex flex-col lg:flex-row justify-between\"><div class=\"space-y-8 md:space-y-16 mb-16 md:mb-0\"><h6 class=\"font-medium text-base uppercase\">Company</h6><p class=\"font-semibold text-base\">Fall Creek <span class=\"font-normal text-gray-300\">/ Variety Royalties</span></p><p class=\"font-semibold text-base\">Source Allies (in-house) <span class=\"font-normal text-gray-300\">/ AMA</span></p><p class=\"font-semibold text-base\">Corteva <span class=\"font-normal text-gray-300\">/ SeedPro</span></p><p class=\"font-semibold text-base\">Principal Financial Group <span class=\"font-normal text-gray-300\">/ ActMod Integrate</span></p><p class=\"font-semibold text-base\">Corteva <span class=\"font-normal text-gray-300\">/ Grain Desk Pro</span></p><p class=\"font-semibold text-base\">AgVision <span class=\"font-normal text-gray-300\">/ Payroll and Financial System</span></p></div><div class=\"space-y-8 md:space-y-16 mb-16 md:mb-0\"><h6 class=\"font-medium text-base uppercase\">Position</h6><p class=\"font-normal text-base\"></p><p class=\"font-normal text-base\">Senior Fullstack Developer / Cloud Architect</p><p class=\"font-normal text-base\">DevOps Engineer</p><p class=\"font-normal text-base\">Tech Lead / Cloud Engineer</p><p class=\"font-normal text-base\">Tech Lead / Cloud Architect</p><p class=\"font-normal text-base\">Senior Software Developer</p><p class=\"font-normal text-base\">Software Developer</p></div><div class=\"space-y-8 md:space-y-16\"><h6 class=\"font-medium text-base uppercase\">Year</h6><p class=\"font-normal text-base\">2023</p><p class=\"font-normal text-base\">2023</p><p class=\"font-normal text-base\">2021</p><p class=\"font-normal text-base\">2021</p><p class=\"font-normal text-base\">2019</p><p class=\"font-normal text-base\">2012</p></div></div></div></section><section class=\"mt-7 md:mt-16\"><div class=\"container max-w-screen-xl mx-auto px-4\"><div class=\"flex flex-col lg:flex-row justify-between\"><div class=\"mb-10 lg:mb-0\"><h1 class=\"font-medium text-3xl md:text-4xl mb-5\">Portfolio</h1><p class=\"font-normal text-xs md:text-base\">Here are my some of my favorite works <br>as a developer.</p></div><div class=\"space-y-24\"><div class=\"flex space-x-6\"><h1 class=\"font-normal text-3xl md:text-4xl\">01</h1><span class=\"w-28 h-0.5 bg-gray-300 mt-5\"></span><div><h1 class=\"font-normal text-3xl md:text-4xl mb-5\">Demo API Generator</h1><p class=\"font-normal text-sm md:text-base\">A dummy data free and documented API generator to facilitate <br>the process of testing the front-end portion of projects.</p></div></div><div class=\"flex space-x-6\"><h1 class=\"font-normal text-3xl md:text-4xl\">02</h1><span class=\"w-28 h-0.5 bg-gray-300 mt-5\"></span><div><h1 class=\"font-normal text-3xl md:text-4xl mb-5\">Demo API Generator</h1><p class=\"font-normal text-sm md:text-base\">A dummy data free and documented API generator to facilitate <br>the process of testing the front-end portion of projects.</p></div></div><div class=\"flex space-x-6\"><h1 class=\"font-normal text-3xl md:text-4xl\">03</h1><span class=\"w-28 h-0.5 bg-gray-300 mt-5\"></span><div><h1 class=\"font-normal text-3xl md:text-4xl mb-5\">Demo API Generator</h1><p class=\"font-normal text-sm md:text-base\">A dummy data free and documented API generator to facilitate <br>the process of testing the front-end portion of projects.</p></div></div></div></div></div></section><section class=\"py-10 md:py-16\"><div class=\"container max-w-screen-xl mx-auto px-4\"><div class=\"text-center\"><p class=\"font-medium text-xs md:text-base\">In my years of experience, I've used C#, python, and Go for backend projects<br>and Angular, NextJS, and React for front-end projects. </p></div></div></section><section class=\"mt-7 md:mt-16\"><div class=\"container max-w-screen-xl mx-auto px-4\"><h1 class=\"font-medium text-3xl md:text-4xl mb-5\">Education</h1><div class=\"flow-root\"><p class=\"float-start font-normal text-xs md:text-base mb-20\">Indian Hills Community College</p><p class=\"float-end font-normal text-xs md:text-base mb-20\">2009 - 2011</p></div></div></section><section class=\"mt-7 md:mt-16\"><div class=\"container max-w-screen-xl mx-auto px-4\"><h1 class=\"font-medium text-3xl md:text-4xl mb-5\">Testimonial</h1><p class=\"font-normal text-xs md:text-base mb-10 md:mb-20\">Below are a few pieces of feedback I have received from colleagues over the years</p><div class=\"grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6\"><div class=\"custom-card\"><p class=\"font-normal text-md mb-4\">Lorem ipsum dolor sit amet, consectetur <br>adipiscing elit, sed do eiusmod tempor <br>incididunt ut labore et dolore magna aliqua.</p><h6 class=\"font-semibold text-md\">Stephan Clark <span class=\"font-medium text-sm\">- CEO at EarlyBird</span></h6></div><div class=\"custom-card\"><p class=\"font-normal text-md mb-4\">Lorem ipsum dolor sit amet, consectetur <br>adipiscing elit, sed do eiusmod tempor <br>incididunt ut labore et dolore magna aliqua.</p><h6 class=\"font-semibold text-md\">Stephan Clark <span class=\"font-medium text-sm\">- CEO at EarlyBird</span></h6></div><div class=\"custom-card\"><p class=\"font-normal text-md mb-4\">Lorem ipsum dolor sit amet, consectetur <br>adipiscing elit, sed do eiusmod tempor <br>incididunt ut labore et dolore magna aliqua.</p><h6 class=\"font-semibold text-md\">Stephan Clark <span class=\"font-medium text-sm\">- CEO at EarlyBird</span></h6></div><div class=\"custom-card\"><p class=\"font-normal text-md mb-4\">Lorem ipsum dolor sit amet, consectetur <br>adipiscing elit, sed do eiusmod tempor <br>incididunt ut labore et dolore magna aliqua.</p><h6 class=\"font-semibold text-md\">Stephan Clark <span class=\"font-medium text-sm\">- CEO at EarlyBird</span></h6></div><div class=\"custom-card\"><p class=\"font-normal text-md mb-4\">Lorem ipsum dolor sit amet, consectetur <br>adipiscing elit, sed do eiusmod tempor <br>incididunt ut labore et dolore magna aliqua.</p><h6 class=\"font-semibold text-md\">Stephan Clark <span class=\"font-medium text-sm\">- CEO at EarlyBird</span></h6></div><div class=\"custom-card\"><p class=\"font-normal text-md mb-4\">Lorem ipsum dolor sit amet, consectetur <br>adipiscing elit, sed do eiusmod tempor <br>incididunt ut labore et dolore magna aliqua.</p><h6 class=\"font-semibold text-md\">Stephan Clark <span class=\"font-medium text-sm\">- CEO at EarlyBird</span></h6></div></div></div></section>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}
