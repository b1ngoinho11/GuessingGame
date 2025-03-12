import { useState, useEffect } from "react";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Card, CardHeader, CardContent, CardFooter } from "@/components/ui/card";
import { Avatar, AvatarImage, AvatarFallback } from "@/components/ui/avatar";
import { DropdownMenu, DropdownMenuTrigger, DropdownMenuContent, DropdownMenuItem } from "@/components/ui/dropdown-menu";
import { Dialog, DialogTrigger, DialogContent, DialogHeader, DialogFooter, DialogTitle, DialogDescription } from "@/components/ui/dialog";
import Cookies from "js-cookie";

export default function App() {
  const [isAuthenticated, setIsAuthenticated] = useState(false);
  const [guess, setGuess] = useState("");
  const [isLoginModalOpen, setIsLoginModalOpen] = useState(false);
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [loginError, setLoginError] = useState("");

  useEffect(() => {
    const token = Cookies.get("authToken");
    if (token) {
      setIsAuthenticated(true);
    }
  }, []);

  const handleAuth = () => {
    if (isAuthenticated) {
      Cookies.remove("authToken");
      setIsAuthenticated(false);
    } else {
      setIsLoginModalOpen(true);
      setLoginError("");
    }
  };

  const handleLogin = () => {
    if (username === "user" && password === "password") {
      Cookies.set("authToken", "your_token_value", {expires: 7});
      setIsAuthenticated(true);
      setIsLoginModalOpen(false);
    } else {
      setLoginError("Invalid username or password. Please try again.");
    }
  };

  const handleGuess = () => {
    if (isAuthenticated) {
      alert(`Your guess: ${guess}`);
    } else {
      alert("You must log in to play the game.");
    }
  };

  return (
    <div className="flex flex-col min-h-screen items-center justify-center bg-neutral-800 text-white dark">
      <div className="absolute top-4 right-4">
        {isAuthenticated ? (
          <DropdownMenu>
            <DropdownMenuTrigger>
              <Avatar>
                <AvatarImage src="/path/to/avatar.jpg" alt="User Avatar" />
                <AvatarFallback>U</AvatarFallback>
              </Avatar>
            </DropdownMenuTrigger>
            <DropdownMenuContent align="end" className="bg-card text-card-foreground">
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
          <Button onClick={handleGuess}>Submit Guess</Button>
        </CardFooter>
      </Card>

      <Dialog open={isLoginModalOpen} onOpenChange={setIsLoginModalOpen}>
        <DialogTrigger />
        <DialogContent className="bg-card text-card-foreground p-8 rounded-lg w-full max-w-md">
          <DialogHeader>
            <DialogTitle>Login</DialogTitle>
            <DialogDescription className={loginError ? "text-red-500" : ""}>
              {loginError ? loginError : "Enter your username and password to log in."}
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
            <Button variant="outline" onClick={() => setIsLoginModalOpen(false)}>
              Cancel
            </Button>
          </DialogFooter>
        </DialogContent>
      </Dialog>
    </div>
  );
}
