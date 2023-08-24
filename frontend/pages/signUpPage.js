import React, { useState, useRef, useEffect } from "react";
import Head from "next/head";
import { FaRegEnvelope, FaTimes } from "react-icons/fa";
import { MdPermIdentity, MdLockOutline } from "react-icons/md";
import {BiSolidInfoCircle, BiHide, BiShow } from "react-icons/bi";
import {ImCheckmark} from "react-icons/im";
import { FcCheckmark } from "react-icons/fc";
import axios from "./api/axios";

// export default function SignUpPage() {

const fullname_valid = /^[A-Z][a-zA-Z\s]*$/;
const username_valid = /^[a-zA-Z][a-zA-Z0-9_]{4,14}$/; //bebas min 5-15 huruf bole spasi
const email_valid = /^[a-zA-Z0-9_.]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/; //dapat menggunakan angka, huruf, _ dan . sebelum @
const pass_valid = /^(?=.*[A-Z])(?=.*[a-z])(?=.*[0-9])(?=.*[^A-Za-z0-9]).{6,10}$/; //ada satu huruf dan angka yang required dengan min.6-10 huruf
const signup_url = "/user/SignUp";

const SignUpPage = () => {
  const userReference = useRef();
  const errReference = useRef();

  // Toggle Password
  const [showPassword, setShowPassword] = useState(false);

  const [fullname, setFullname] = useState("");
  const [validFullname, setValidFullname] = useState(false);
  const [fullnameFocus, setFullnameFocus] = useState(false);

  const [uname, setUname] = useState("");
  const [validUname, setValidUname] = useState(false);
  const [unameFocus, setUnameFocus] = useState(false);

  const [email, setEmail] = useState("");
  const [validEmail, setValidEmail] = useState(false);
  const [emailFocus, setEmailFocus] = useState(false);

  const [pass, setPass] = useState("");
  const [validPass, setValidPass] = useState(false);
  const [passFocus, setPassFocus] = useState(false);

  // Error Message
  const [errMsg, setErrMsg] = useState("");
  const [success, setSuccess] = useState(false);

  useEffect(() => {
    userReference.current.focus(); //setting focusnya ketika komponen load
  }, []);

  useEffect(() => {
    const result = fullname_valid.test(fullname);
    console.log(result);
    console.log(fullname);
    setValidFullname(result);
  }, [fullname]);

  useEffect(() => {
    const result = username_valid.test(uname);
    console.log(result);
    console.log(uname);
    setValidUname(result);
  }, [uname]);

  useEffect(() => {
    const result = email_valid.test(email);
    console.log(result);
    console.log(email);
    setValidEmail(result);
  }, [email]);

  useEffect(() => {
    const result = pass_valid.test(pass);
    console.log(result);
    console.log(pass);
    setValidPass(result);
  }, [pass]);

  useEffect(() => {
    setErrMsg("");
  }, [fullname, uname, email, pass]);

  const handleSubmit = async (e) => {
    e.preventDefault();

    try {
      const data = {
        fullname,
        username: uname,
        email,
        password: pass,
      };

      axios
        .post(signup_url, data)
        .then((res) => {
          console.log("success");
        })
        .catch((err) => {
          console.log(err);
        });

    } catch (err) {
      console.log('Registrastion Failed',err)
      errReference.current.focus();
    }
    setSuccess(true);
  };

  return (
    <div>
      <div>
        <Head>
          <title>Sign Up</title>
        </Head>
      </div>
      {/* Backkground */}
      <div className="bg-gradient-to-r from-cyan-500/10 via-teal-300/10 to-sky-200/10 block h-screen items-center justify-center p-4 md:flex">
        {/* Container */}
        <div className=" bg-white flex flex-col items-center max-w-screen-lg overflow-hidden rounded-lg shadow-[0_3px_10px_rgb(0,0,0,1)] space-y-8 w-full md:flex-row text-slate-700 ">
          {/* Welcome*/}
          <div className="bg-gradient-to-t from-sky-200 to-stone-50 flex flex-col items-center p-4 w-full md:w-2/5 text-white">
            <div className="flex flex-col justify-center items-start text-left text-slate-700 w-full p-10">
              <h2 className="text-3xl font-bold mb-4 ">Welcome To Jico</h2>
              <p className="text-base font-normal mb-48">
                Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nam nec
                ultricies nisi. Suspendisse pulvinar viverra nibh vel ultricies.
                Mauris tincidunt mollis diam, at mollis enim aliquet eget. Fusce
                eros neque, pharetra eget tincidunt in, tincidunt nec tellus.
              </p>

              <div className="h-8 w-16 flex object-right-bottom">
                <img
                  className=""
                  src="https://upload.wikimedia.org/wikipedia/id/thumb/c/c4/Telkom_Indonesia_2013.svg/1200px-Telkom_Indonesia_2013.svg.png"
                />
                <img
                  className=""
                  src="https://cdn.icon-icons.com/icons2/2699/PNG/512/atlassian_jira_logo_icon_170512.png"
                />
              </div>
            </div>
          </div>

     
          {/* Sign Up */}
          {/* Form */}
          <p
            ref={errReference}
            className={` ${errMsg ? "errmsg" : "offscreen"}`}
            aria-live="assertive"
          >
            {errMsg}
          </p>

          <form
            onSubmit={handleSubmit}
            className="flex flex-col justify-center items-center p-14 pl-36 space-y-4"
          >
            {/* Title */}
            <div className="flex flex-col  items-center space-y-4 mb-4  ">
              <h1 className="text-slate-700 text-4xl font-semibold ">
                Sign Up
              </h1>
            </div>
            <div className="relative py-2">

              {/* Fullname */}
                {/* Warning Fullname*/}
                <p
                  id="note"
                  className={`${
                    fullnameFocus && fullname && !validFullname 
                      ? "instructions"
                      : "sr-only"
                    } flex items-center text-red-600 mb-2`
                  } 
                >
                  <BiSolidInfoCircle className="mr-2" /> 
                  Must start with a capital letter
                </p>
                
              <div className="w-full bg-white flex items-center mb-[3%] border-gray-300 border rounded-lg px-3 py-2 focus:outline-none shadow shadow-black ">
                <MdPermIdentity className="m-[1%] text-slate-700" />
                <input
                  className="pl-2 py-1 w-72 focus:outline-none"
                  type="text"
                  name="name"
                  placeholder="Full Name"
                  ref={userReference}
                  autoComplete="off"
                  onChange={(e) => setFullname(e.target.value)}
                  required
                  aria-invalid={validFullname ? "false" : "true"}
                  aria-describedby="note"
                  onFocus={() => setFullnameFocus(true)}
                  onBlur={() => setFullnameFocus(false)}
                />

                {/* checkmark */}
                <span className={validFullname ? "valid" : "hidden"}>
                  <FcCheckmark />
                </span>
                <span
                  className={`${
                    validFullname || !fullname ? "hidden" : ""
                  } text-red-600`}
                >
                  <FaTimes />
                </span>
              </div>

              {/* Username */}
                 {/* Warning  Username*/}
                 <p 
                  id="uname-note"
                  className={`${
                    unameFocus && uname && !validUname
                      ? "instructions"
                      : "sr-only"
                    } flex items-center text-red-600 mt-2`
                  } 
                >
                  <BiSolidInfoCircle  className="mr-2"/>
                  Must be 5 to 15 characters, letters, numbers are allowed
                </p>

              <div className="w-full bg-white flex items-center mb-[3%] border-gray-300 border rounded-lg px-3 py-2 focus:outline-none shadow shadow-black ">
                <MdPermIdentity className="m-[1%] text-slate-700" />
                <input
                  className="pl-2 py-1 w-72 focus:outline-none"
                  type="text"
                  name="username"
                  placeholder="Username"
                  autoComplete="off"
                  onChange={(e) => setUname(e.target.value)}
                  required
                  aria-invalid={validUname ? "false" : "true"}
                  aria-describedby="uname-note"
                  onFocus={() => setUnameFocus(true)}
                  onBlur={() => setUnameFocus(false)}
                />

                {/* checkmark */}
                <span className={validUname ? "valid" : "hidden"}>
                  <FcCheckmark />
                </span>
                <span
                  className={`${
                    validUname || !uname ? "hidden" : ""
                  } text-red-600`}
                >
                  <FaTimes />
                </span>
              </div>

              {/* Email */}
                {/* Warning */}
                  <p
                  id="email-note"
                  className={`${
                    emailFocus && email && !validEmail
                      ? "instructions"
                      : "sr-only"
                    } flex items-center text-red-600 mt-2`
                  }
                >
                  <BiSolidInfoCircle className="mr-2"/>               
                  Letters, numbers, underscore and dot are allowed
                </p>

              <div className=" bg-white flex items-center mb-[3%] border-gray-300 border rounded-lg px-3 py-2 focus:outline-none shadow shadow-black ">
                <FaRegEnvelope className="m-[1%] text-slate-700" />
                <input
                  className="pl-2 py-1 w-72 focus:outline-none "
                  type="email"
                  name="email"
                  placeholder="Email"
                  onChange={(e) => setEmail(e.target.value)}
                  required
                  aria-invalid={validEmail ? "false" : "true"}
                  aria-describedby="email-note"
                  onFocus={() => setEmailFocus(true)}
                  onBlur={() => setEmailFocus(false)}
                />

                {/* checkmark */}
                <span className={validEmail ? "valid" : "hidden"}>
                  <FcCheckmark />
                </span>
                <span
                  className={`${
                    validEmail || !email ? "hidden" : ""
                  } text-red-600`}
                >
                  <FaTimes />
                </span>
              </div>

              {/*Password*/}
                {/* Warning */}
                 <p
                  id="email-note"
                  className={`${
                    passFocus && pass && !validPass 
                    ? "instructions" 
                    : "sr-only"
                    } flex items-center text-red-600 mt-2`
                  }
                >
                  <BiSolidInfoCircle className="mr-2"/>
                  6 to 10 characters <br />
                  Must have atleast 1 characters, 1 numbers
                </p>

              <div className="w-full bg-white flex items-center mb-[3%] border-gray-300 border rounded-lg px-3 py-2 focus:outline-none shadow shadow-black ">
                <MdLockOutline className="m-[1%] text-slate-700" />
                <input
                  className="pl-2 py-1 w-72 focus:outline-none"
                  type={showPassword ? "text" : "password"}
                  name="password"
                  placeholder="Password"
                  onChange={(e) => setPass(e.target.value)}
                  required
                  aria-invalid={validPass ? "false" : "true"}
                  aria-describedby="pass-note"
                  onFocus={() => setPassFocus(true)}
                  onBlur={() => setPassFocus(false)}
                />
                {showPassword ? (
                  <BiHide
                    onClick={() => setShowPassword(false)}
                    className="cursor-pointer"
                  />
                ) : (
                  <BiShow
                    onClick={() => setShowPassword(true)}
                    className="cursor-pointer"
                  />
                )}

                {/* checkmark */}
                <span className={validPass ? "valid" : "hidden"}>
                  <FcCheckmark />
                </span>
                <span
                  className={`${
                    validPass || !pass ? "hidden" : ""
                  } text-red-600`}
                >
                  <FaTimes />
                </span>
              </div>

              {/* Button Submit */}
              <div className="text-center mb-6">
                <button
                  type="submit"
                  className="bg-blue-500 text-white w-80 px-4 py-2 rounded-lg hover:bg-blue-600 transition-colors duration-200"
                >
                  Sign Up
                </button>
              </div>

              {/* Sign in */}
              <div className="flex flex-col items-center mt-3 leading-snug">
                <p>
                  Already a member?
                  <a
                    href="/signInPage"
                    className="text-blue-700 font-bold text-base leading-snug hover:underline "
                  >
                    Sign In
                  </a>
                </p>
              </div>
            </div>
          </form>

          {/* pop up succes */}
          {success && (
            <div className="fixed inset-0 flex items-center justify-center z-50 bg-black bg-opacity-70">
              <div className="bg-white p-8 rounded-lg flex flex-col text-white ">
                <h2 className="text-xl font-semibold mb-4 text-black text-center">
                  Successfully Sign Up
                </h2>
                <ImCheckmark className="text-6xl text-sky-700 animate-pulse mx-auto mb-4" />
                <a href='/signInPage' className="text-sky-700 text-center text-lg font-semibold hover:underline ">Sign in</a>
                <button
                  onClick={() => setSuccess(false)}
                  className="bg-blue-500 px-4 py-2 hover:bg-blue-700 font-bold rounded-lg mt-4 text-white"
                >
                  Close
                </button>
              </div>
            </div>
          )}
        </div>
      </div>
    </div>
  );
};

export default SignUpPage;
