require 'ffi'

module ArrayStats
  extend FFI::Library

  ffi_lib './libgostats.so'
  attach_function :percentile_go, [:pointer, :double], :double
end

class Array
  include ArrayStats

  def percentile(percent)
    # Create the pointer to the array
    pointer = FFI::MemoryPointer.new :double, self.size

    # Fill the memory location with your data
    pointer.put_array_of_double 0, self

    percentile_go(pointer, percent)
  end
end

p [1,2,3,4,5,6].percentile(25)
