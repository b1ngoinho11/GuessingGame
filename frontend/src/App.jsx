import { useState, useEffect } from "react";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import {
  Card,
  CardHeader,
  CardContent,
  CardFooter,
  CardDescription,
} from "@/components/ui/card";
import { Avatar, AvatarImage, AvatarFallback } from "@/components/ui/avatar";
import {
  DropdownMenu,
  DropdownMenuTrigger,
  DropdownMenuContent,
  DropdownMenuItem,
} from "@/components/ui/dropdown-menu";
import {
  Dialog,
  DialogTrigger,
  DialogContent,
  DialogHeader,
  DialogFooter,
  DialogTitle,
  DialogDescription,
} from "@/components/ui/dialog";
import Cookies from "js-cookie";

export default function App() {
  const [isAuthenticated, setIsAuthenticated] = useState(false);
  const [guess, setGuess] = useState("");
  const [isLoginModalOpen, setIsLoginModalOpen] = useState(false);
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [loginError, setLoginError] = useState("");
  const [messageModalOpen, setMessageModalOpen] = useState(false);
  const [messageContent, setMessageContent] = useState("");

  useEffect(() => {
    const token = Cookies.get("token");
    if (token) {
      setIsAuthenticated(true);
    }
  }, []);

  const showMessageModal = (message) => {
    setMessageContent(message);
    setMessageModalOpen(true);
  };

  const handleAuth = () => {
    if (isAuthenticated) {
      Cookies.remove("token");
      setIsAuthenticated(false);
    } else {
      setIsLoginModalOpen(true);
      setLoginError("");
    }
  };

  const handleLogin = async () => {
    try {
      const response = await fetch("http://127.0.0.1:3000/login", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          username: username,
          password: password,
        }),
        credentials: "include",
      });
  
      if (response.ok) {
        setIsAuthenticated(true);
        setIsLoginModalOpen(false);
        setUsername("");
        setPassword("");
      } else {
        const errorData = await response.json();
        setLoginError("Login failed. Please try again");
      }
    } catch (error) {
      setLoginError("An error occurred. Please try again.");
      console.error("Login error:", error);
    }
  };

  const handleGuess = async () => {
    if (isAuthenticated) {
      try {
        const token = Cookies.get("token");
        if (!token) {
          showMessageModal("Session expired. Please log in again.");
          setIsAuthenticated(false);
          return;
        }
        
        const response = await fetch(`http://127.0.0.1:3000/guess/${guess}`, {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          credentials: "include",
        });
        
        if (response.ok) {
          const data = await response.json();
          showMessageModal(data.message);
        } else {
          const errorData = await response.json();
          
          if (response.status === 401) {
            showMessageModal("Your session has expired. Please log in again.");
            setIsAuthenticated(false);
          } else {
            showMessageModal(errorData.error || "Something went wrong.");
          }
        }
      } catch (error) {
        showMessageModal("An error occurred. Please try again.");
        console.error("Guess error:", error);
      }
    } else {
      showMessageModal("You must log in to play the game.");
    }
  };

  return (
    <div className="flex flex-col min-h-screen items-center justify-center bg-neutral-800 text-white dark">
      <div className="absolute top-4 right-4">
        {isAuthenticated ? (
          <DropdownMenu>
            <DropdownMenuTrigger>
              <Avatar>
                <AvatarImage src="https://github.com/shadcn.png" />
                <AvatarFallback>CN</AvatarFallback>
              </Avatar>
            </DropdownMenuTrigger>
            <DropdownMenuContent
              align="end"
              className="bg-card text-card-foreground"
            >
              <DropdownMenuItem onClick={handleAuth}>Logout</DropdownMenuItem>
            </DropdownMenuContent>
          </DropdownMenu>
        ) : (
          <Button onClick={handleAuth}>Login</Button>
        )}
      </div>

      <Card className="w-full max-w-md shadow-lg bg-card text-card-foreground">
        <CardHeader>
          <h2 className="text-xl font-semibold text-center">Guessing Game</h2>
        </CardHeader>
        <CardDescription>
          <h3 className="text-l text-center"> Guess a number 0-9</h3>
        </CardDescription>
        <CardContent className="flex flex-col gap-4">
          <Input
            type="text"
            placeholder="Enter your guess..."
            value={guess}
            onChange={(e) => setGuess(e.target.value)}
            className="bg-input text-input-foreground"
          />
        </CardContent>
        <CardFooter className="flex justify-center">
          <Button onClick={handleGuess}>Submit</Button>
        </CardFooter>
      </Card>

      <Dialog open={isLoginModalOpen} onOpenChange={setIsLoginModalOpen}>
        <DialogTrigger />
        <DialogContent className="bg-card text-card-foreground p-8 rounded-lg w-full max-w-md">
          <DialogHeader>
            <DialogTitle>Login</DialogTitle>
            <DialogDescription className={loginError ? "text-red-500" : ""}>
              {loginError
                ? loginError
                : "Enter your username and password to log in."}
            </DialogDescription>
          </DialogHeader>
          <div className="flex flex-col gap-4">
            <Input
              type="text"
              placeholder="Username"
              value={username}
              onChange={(e) => setUsername(e.target.value)}
              className="bg-input text-input-foreground"
            />
            <Input
              type="password"
              placeholder="Password"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              className="bg-input text-input-foreground"
            />
          </div>
          <DialogFooter>
            <Button onClick={handleLogin}>Login</Button>
            <Button
              variant="outline"
              onClick={() => setIsLoginModalOpen(false)}
            >
              Cancel
            </Button>
          </DialogFooter>
        </DialogContent>
      </Dialog>

      <Dialog open={messageModalOpen} onOpenChange={setMessageModalOpen}>
        <DialogContent className="bg-card text-card-foreground p-8 rounded-lg w-full max-w-md">
          <DialogHeader>
            <DialogTitle>Alert</DialogTitle>
            <DialogDescription>{messageContent}</DialogDescription>
          </DialogHeader>
          <DialogFooter>
            <Button onClick={() => setMessageModalOpen(false)}>OK</Button>
          </DialogFooter>
        </DialogContent>
      </Dialog>
    </div>
  );
}