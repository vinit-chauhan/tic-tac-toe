import axios from "axios";
import * as SecureStore from "expo-secure-store";
import { createContext, useContext, useEffect, useState } from "react";
import { Platform } from "react-native";

interface AuthProps {
  authState?: { token: string | null; authenticated: boolean | null };
  onRegister?: (
    username: string,
    password: string,
    email: string | null
  ) => Promise<any>;
  onLogin?: (username: string, password: string) => Promise<any>;
  onLogout?: () => Promise<any>;
}

const TOKEN_KEY = "jwt-token";
export const API_URL = process.env.BACKEND_API_URL || "http://192.168.140.10:3000";
const AuthContext = createContext<AuthProps>({});

export const useAuth = () => {
  return useContext(AuthContext);
};

type AuthStateType = {
  token: string | null;
  authenticated: boolean | null;
};

export const AuthProvider = ({ children }: any) => {
  const [authState, setAuthState] = useState<AuthStateType>({
    token: null,
    authenticated: null,
  });

  useEffect(() => {
    const loadToken = async () => {
      let token;

      if (Platform.OS == "web") {
        token = localStorage.getItem(TOKEN_KEY);
      } else {
        token = await SecureStore.getItemAsync(TOKEN_KEY);
      }

      if (token) {
        axios.defaults.headers.common["Authorization"] = `Bearer ${token}`;
        setAuthState({
          token: token,
          authenticated: true,
        });
      }
    };

    loadToken();
  }, []);

  const register = async (
    username: string,
    password: string,
    email: string | null
  ) => {
    try {
      return await axios.post(
        `${API_URL}/users`,
        { username, password, email },
        {
          headers: {
            "Content-Type": "application/json",
          },
        }
      );
    } catch (e) {
      return { error: true, msg: (e as any).response.data.msg };
    }
  };

  const login = async (username: string, password: string) => {
    try {
      const result = await axios({
        method: "post",
        url: `${API_URL}/auth/login`,
        data: { username, password },
        headers: {
          "Content-Type": "application/json",
        },
      });

      axios.defaults.headers.common[
        "Authorization"
      ] = `Bearer ${result.data.token}`;

      setAuthState({
        token: result.data.token,
        authenticated: true,
      });

      if (Platform.OS == "web") {
        localStorage.setItem(TOKEN_KEY, result.data.token);
      } else {
        await SecureStore.setItemAsync(TOKEN_KEY, result.data.token);
      }

      return result;
    } catch (e) {
      return { error: true, msg: (e as any).response?.data?.msg };
    }
  };

  const logout = async () => {
    if (Platform.OS == "web") {
      localStorage.removeItem(TOKEN_KEY);
    } else {
      await SecureStore.deleteItemAsync(TOKEN_KEY);
    }

    axios.defaults.headers.common["Authorization"] = "";
    setAuthState({
      token: null,
      authenticated: false,
    });
  };

  const value = {
    authState: authState,
    onRegister: register,
    onLogin: login,
    onLogout: logout,
  };
  return <AuthContext.Provider value={value}>{children}</AuthContext.Provider>;
};
