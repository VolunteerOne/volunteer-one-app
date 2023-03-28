import React, { useState, useEffect } from 'react';
import { Pressable, Button, Image, View, Platform, StyleSheet } from 'react-native';
import * as ImagePicker from 'expo-image-picker';
import { Text } from 'galio-framework';


export default function ImagePickerExample() {
  const [image, setImage] = useState(null);

  const pickImage = async () => {
    // No permissions request is necessary for launching the image library
    let result = await ImagePicker.launchImageLibraryAsync({
      mediaTypes: ImagePicker.MediaTypeOptions.All,
      allowsEditing: true,
      aspect: [4, 3],
      quality: 1,
    });

    console.log(result);

    if (!result.canceled) {
      setImage(result.assets[0].uri);
    }
  };

  return (
    <View style={{ }}>
      <Pressable
        style={[styles.imageButton]}
        onPress={pickImage}>
      {/* <Button title="Upload Image" onPress={pickImage} style={styles.imageButton}/> */}
      <Text style={styles.textStyle}>UPLOAD IMAGE</Text>
      {image && <Image source={{ uri: image }} style={styles.imagePreview} />}
      </Pressable>
    </View>
  );
}

const styles = StyleSheet.create({
    imageButton: {
    borderRadius: 5,
    padding: 5,
    elevation: 2,
    backgroundColor:'lightgrey',
    alignItems: 'center',
    justifyContent: 'center'
  },

  textStyle: {
    color: 'white',
    fontWeight: 'bold',
    textAlign: 'center',
  },
  imagePreview: { 
    width: 200,
    height: 200,
    borderColor: 'grey',
    borderWidth: .5,
    marginBottom: 5,
    
  },
})