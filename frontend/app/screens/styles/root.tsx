import { StyleSheet } from "react-native";

export const style = StyleSheet.create({
  container: {
    flex: 1,
    justifyContent: "center",
    alignItems: "center",
  },
  textBox: {
    borderRadius: 8,
    borderColor: "#F35369AA",
    borderWidth: 2,
    height: 50,
    width: "80%",
    justifyContent: "center",
    alignItems: "center",
    padding: 10,
    margin: 10,
  },
  link: {
    margin: 15,
    textAlign: "center",
    color: "#00f",
  },
  loginBtn: {
    width: "40%",
    borderRadius: 8,
    height: 50,
    alignItems: "center",
    justifyContent: "center",
    marginTop: 10,
    backgroundColor: "#F35369",
  },
  text: {
    color: "white",
    fontFamily: "Arial",
    fontWeight: "bold",
    fontSize: 16,
  },
  backdrop: {
    backgroundColor: "#FAFFD2",
  },
  formGroup: {
    width: "50%",
  },
});
